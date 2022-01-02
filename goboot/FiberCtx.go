package goboot

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
	"github.com/lamgor666/goboot-common/AppConf"
	"github.com/lamgor666/goboot-common/enum/RegexConst"
	"github.com/lamgor666/goboot-common/util/castx"
	"github.com/lamgor666/goboot-common/util/jsonx"
	"github.com/lamgor666/goboot-common/util/mapx"
	"github.com/lamgor666/goboot-common/util/numberx"
	"github.com/lamgor666/goboot-common/util/slicex"
	"github.com/lamgor666/goboot-common/util/stringx"
	"github.com/lamgor666/goboot-core/enum/ReqParamSecurityMode"
	"math"
	"mime/multipart"
	"net/url"
	"regexp"
	"strings"
)

func fiberGetMethod(ctx *fiber.Ctx) string {
	return ctx.Method()
}

func fiberGetHeader(ctx *fiber.Ctx, name string) string {
	return ctx.Get(name)
}

func fiberGetHeaders(ctx *fiber.Ctx) map[string]string {
	map1 := map[string]string{}

	ctx.Request().Header.VisitAll(func(keyBytes, valueBytes []byte) {
		var key string

		if len(keyBytes) > 0 {
			key = string(utils.CopyBytes(keyBytes))
		}

		if key == "" {
			return
		}

		var value string

		if len(valueBytes) > 0 {
			value = string(utils.CopyBytes(valueBytes))
		}

		map1[key] = value
	})

	return map1
}

func fiberGetQueryParams(ctx *fiber.Ctx) map[string]string {
	s1 := utils.CopyString(ctx.OriginalURL())

	if !strings.Contains(s1, "?") {
		return map[string]string{}
	}

	values, err := url.ParseQuery(stringx.SubstringAfter(s1, "?"))

	if err != nil || len(values) < 1 {
		return map[string]string{}
	}

	map1 := map[string]string{}
	
	for key, parts := range values {
		if key == "" {
			continue
		}

		if len(parts) > 0 {
			map1[key] = parts[0]
		} else {
			map1[key] = ""
		}
	}
	
	return map1
}

func fiberGetQueryString(ctx *fiber.Ctx, urlencode ...bool) string {
	if len(urlencode) < 1 || !urlencode[0] {
		s1 := utils.CopyString(ctx.OriginalURL())

		if !strings.Contains(s1, "?") {
			return ""
		}

		return stringx.SubstringAfter(s1, "?")
	}
	
	map1 := fiberGetQueryParams(ctx)
	
	if len(map1) < 1 {
		return ""
	}

	values := url.Values{}

	for key, value := range map1 {
		values[key] = []string{value}
	}

	return values.Encode()
}

func fiberGetRequestUrl(ctx *fiber.Ctx, withQueryString ...bool) string {
	s1 := utils.CopyString(ctx.OriginalURL())
	
	if s1 == "" {
		s1 = "/"
	} else {
		s1 = stringx.EnsureLeft(s1, "/")
	}
	
	if len(withQueryString) < 1 || !withQueryString[0] {
		return s1
	}
	
	if strings.Contains(s1, "?") {
		s1 = stringx.SubstringBefore(s1, "?")
	}
	
	return s1
}

func fiberGetFormData(ctx *fiber.Ctx) map[string]string {
	map1 := map[string]string{}
	isPost := ctx.Request().Header.IsPost()

	if !isPost {
		return map1
	}

	contentType := strings.ToLower(ctx.Get(fiber.HeaderContentType))
	isPostForm := strings.Contains(contentType, fiber.MIMEApplicationForm)
	isMultipartForm := strings.Contains(contentType, fiber.MIMEMultipartForm)

	if isPostForm {
		ctx.Request().PostArgs().VisitAll(func(keyBytes, valueBytes []byte) {
			var key string

			if len(keyBytes) > 0 {
				key = string(utils.CopyBytes(keyBytes))
			}

			if key == "" {
				return
			}

			var value string

			if len(valueBytes) > 0 {
				value = string(utils.CopyBytes(valueBytes))
			}

			map1[key] = value
		})

		return map1
	}

	if isMultipartForm {
		form, err := ctx.MultipartForm()

		if err != nil {
			return map1
		}

		for key, values := range form.Value {
			if key == "" {
				continue
			}

			if len(values) > 0 {
				map1[key] = values[0]
			} else {
				map1[key] = ""
			}
		}

		return map1
	}

	return map1
}

func fiberGetClientIp(ctx *fiber.Ctx) string {
	ips := ctx.IPs()

	if len(ips) > 0 {
		for _, s1 := range ips {
			if stringx.RegexMatch(s1, "^[0-9.]+$") {
				return s1
			}
		}
	}

	ip := ctx.Get("X-Real-IP")

	if ip == "" {
		ip = ctx.IP()
	}

	parts := stringx.SplitWithRegexp(strings.TrimSpace(ip), RegexConst.CommaSep)

	if len(parts) < 1 {
		return ""
	}

	return strings.TrimSpace(parts[0])
}

func fiberPathvariable(ctx *fiber.Ctx, name string, defaultValue ...interface{}) string {
	var dv string

	if len(defaultValue) > 0 {
		if s1, err := castx.ToStringE(defaultValue[0]); err == nil {
			dv = s1
		}
	}

	return ctx.Params(name, dv)
}

func fiberPathvariableBool(ctx *fiber.Ctx, name string, defaultValue ...interface{}) bool {
	var dv bool

	if len(defaultValue) > 0 {
		if b1, err := castx.ToBoolE(defaultValue[0]); err == nil {
			dv = b1
		}
	}

	if b1, err := castx.ToBoolE(ctx.Params(name)); err == nil {
		return b1
	}

	return dv
}

func fiberPathvariableInt(ctx *fiber.Ctx, name string, defaultValue ...interface{}) int {
	dv := math.MinInt32

	if len(defaultValue) > 0 {
		if n1, err := castx.ToIntE(defaultValue[0]); err == nil {
			dv = n1
		}
	}

	return castx.ToInt(ctx.Params(name), dv)
}

func fiberPathvariableInt64(ctx *fiber.Ctx, name string, defaultValue ...interface{}) int64 {
	dv := int64(math.MinInt64)

	if len(defaultValue) > 0 {
		if n1, err := castx.ToInt64E(defaultValue[0]); err == nil {
			dv = n1
		}
	}

	return castx.ToInt64(ctx.Params(name), dv)
}

func fiberPathvariableFloat32(ctx *fiber.Ctx, name string, defaultValue ...interface{}) float32 {
	dv := float32(math.SmallestNonzeroFloat32)

	if len(defaultValue) > 0 {
		if n1, err := castx.ToFloat32E(defaultValue[0]); err == nil {
			dv = n1
		}
	}

	return castx.ToFloat32(ctx.Params(name), dv)
}

func fiberPathvariableFloat64(ctx *fiber.Ctx, name string, defaultValue ...interface{}) float64 {
	dv := math.SmallestNonzeroFloat64

	if len(defaultValue) > 0 {
		if n1, err := castx.ToFloat64E(defaultValue[0]); err == nil {
			dv = n1
		}
	}

	return castx.ToFloat64(ctx.Params(name), dv)
}

func fiberReqParam(ctx *fiber.Ctx, name string, mode int, defaultValue ...interface{}) string {
	var dv string

	if len(defaultValue) > 0 {
		if s1, err := castx.ToStringE(defaultValue[0]); err == nil {
			dv = s1
		}
	}

	modes := []int{
		ReqParamSecurityMode.None,
		ReqParamSecurityMode.HtmlPurify,
		ReqParamSecurityMode.StripTags,
	}

	if !slicex.InIntSlice(mode, modes) {
		mode = ReqParamSecurityMode.StripTags
	}

	value := fiberFormParam(ctx, name)

	if value == "" {
		value = ctx.Query(name)
	}

	if value == "" {
		return dv
	}

	if mode != ReqParamSecurityMode.None {
		value = stringx.StripTags(value)
	}

	return value
}

func fiberReqParamBool(ctx *fiber.Ctx, name string, defaultValue ...interface{}) bool {
	var dv bool

	if len(defaultValue) > 0 {
		if b1, err := castx.ToBoolE(defaultValue[0]); err == nil {
			dv = b1
		}
	}

	s1 := fiberReqParam(ctx, name, ReqParamSecurityMode.None)

	if b1, err := castx.ToBoolE(s1); err == nil {
		return b1
	}

	return dv
}

func fiberReqParamInt(ctx *fiber.Ctx, name string, defaultValue ...interface{}) int {
	dv := math.MinInt32

	if len(defaultValue) > 0 {
		if n1, err := castx.ToIntE(defaultValue[0]); err == nil {
			dv = n1
		}
	}

	s1 := fiberReqParam(ctx, name, ReqParamSecurityMode.None)
	return castx.ToInt(s1, dv)
}

func fiberReqParamInt64(ctx *fiber.Ctx, name string, defaultValue ...interface{}) int64 {
	dv := int64(math.MinInt64)

	if len(defaultValue) > 0 {
		if n1, err := castx.ToInt64E(defaultValue[0]); err == nil {
			dv = n1
		}
	}

	s1 := fiberReqParam(ctx, name, ReqParamSecurityMode.None)
	return castx.ToInt64(s1, dv)
}

func fiberReqParamFloat32(ctx *fiber.Ctx, name string, defaultValue ...interface{}) float32 {
	dv := float32(math.SmallestNonzeroFloat32)

	if len(defaultValue) > 0 {
		if n1, err := castx.ToFloat32E(defaultValue[0]); err == nil {
			dv = n1
		}
	}

	s1 := fiberReqParam(ctx, name, ReqParamSecurityMode.None)
	return castx.ToFloat32(s1, dv)
}

func fiberReqParamFloat64(ctx *fiber.Ctx, name string, defaultValue ...interface{}) float64 {
	dv := math.SmallestNonzeroFloat64

	if len(defaultValue) > 0 {
		if n1, err := castx.ToFloat64E(defaultValue[0]); err == nil {
			dv = n1
		}
	}

	s1 := fiberReqParam(ctx, name, ReqParamSecurityMode.None)
	return castx.ToFloat64(s1, dv)
}

func fiberGetJwt(ctx *fiber.Ctx) *jwt.Token {
	token := strings.TrimSpace(ctx.Get(fiber.HeaderAuthorization))
	token = stringx.RegexReplace(token, `[\x20\t]+`, " ")

	if strings.Contains(token, " ") {
		token = stringx.SubstringAfter(token, " ")
	}

	if token == "" {
		return nil
	}

	tk, _ := ParseJsonWebToken(token)
	return tk
}

func fiberGetRawBody(ctx *fiber.Ctx) []byte {
	isPost := ctx.Request().Header.IsPost()
	isPut := ctx.Request().Header.IsPut()
	isPatch := ctx.Request().Header.IsPatch()
	isDelete := ctx.Request().Header.IsDelete()
	isJson := (isPost || isPut || isPatch || isDelete) && ctx.Is("json")
	isXml := (isPost || isPut || isPatch || isDelete) && ctx.Is("xml")

	if isJson || isXml {
		if len(ctx.Body()) < 1 {
			return make([]byte, 0)
		}

		buf := utils.CopyBytes(ctx.Body())

		if AppConf.GetBoolean("logging.logGetRawBody") {
			RuntimeLogger().Debug("raw body: " + string(buf))
		}

		return buf
	}

	contentType := strings.ToLower(ctx.Get(fiber.HeaderContentType))
	isPostForm := strings.Contains(contentType, fiber.MIMEApplicationForm)
	isMultipartForm := strings.Contains(contentType, fiber.MIMEMultipartForm)

	if !isPost {
		return make([]byte, 0)
	}

	if !isPostForm && !isMultipartForm {
		return make([]byte, 0)
	}

	formData := fiberGetFormData(ctx)

	if len(formData) < 1 {
		return make([]byte, 0)
	}

	sb := strings.Builder{}

	for key, value := range formData {
		sb.WriteString("&")
		sb.WriteString(key)
		sb.WriteString("=")
		sb.WriteString(value)
	}

	contents := sb.String()

	if contents != "" {
		contents = contents[1:]
	}

	if AppConf.GetBoolean("logging.logGetRawBody") {
		RuntimeLogger().Debug("raw body via form data: " + contents)
	}

	return []byte(contents)
}

func fiberGetMap(ctx *fiber.Ctx, rules ...interface{}) map[string]interface{} {
	var _rules []string

	if len(rules) > 0 {
		if a1, ok := rules[0].([]string); ok && len(a1) > 0 {
			_rules = a1
		} else if s1, ok := rules[0].(string); ok && s1 != "" {
			re := regexp.MustCompile(RegexConst.CommaSep)
			_rules = re.Split(s1, -1)
		}
	}

	isGet := ctx.Request().Header.IsGet()
	isPost := ctx.Request().Header.IsPost()
	isPut := ctx.Request().Header.IsPut()
	isPatch := ctx.Request().Header.IsPatch()
	isDelete := ctx.Request().Header.IsDelete()
	isJson := (isPost || isPut || isPatch || isDelete) && ctx.Is("json")

	if isJson {
		map1 := map[string]interface{}{}

		if len(ctx.Body()) > 0 {
			map1 = jsonx.MapFrom(utils.CopyBytes(ctx.Body()))
		}

		if len(_rules) < 1 {
			return map1
		}

		return fiberGetMapWithRules(map1, _rules)
	}

	isXml := (isPost || isPut || isPatch || isDelete) && ctx.Is("xml")

	if isXml {
		map1 := map[string]string{}

		if len(ctx.Body()) > 0 {
			map1 = mapx.FromXml(utils.CopyBytes(ctx.Body()))
		}

		if len(_rules) < 1 {
			return castx.ToStringMap(map1)
		}

		return fiberGetMapWithRules(castx.ToStringMap(map1), _rules)
	}

	if isGet {
		map1 := fiberGetQueryParams(ctx)

		if len(_rules) < 1 {
			return castx.ToStringMap(map1)
		}

		return fiberGetMapWithRules(castx.ToStringMap(map1), _rules)
	}

	if !isPost {
		return map[string]interface{}{}
	}

	contentType := strings.ToLower(ctx.Get(fiber.HeaderContentType))
	isPostForm := strings.Contains(contentType, fiber.MIMEApplicationForm)
	isMultipartForm := strings.Contains(contentType, fiber.MIMEMultipartForm)

	if !isPostForm && !isMultipartForm {
		return map[string]interface{}{}
	}

	if len(_rules) > 0 {
		return fiberGetMapWithRules(ctx, _rules)
	}

	map1 := map[string]interface{}{}
	queryParams := fiberGetQueryParams(ctx)

	for key, value := range queryParams {
		map1[key] = value
	}

	formData := fiberGetFormData(ctx)

	for key, value := range formData {
		map1[key] = value
	}

	return map1
}

func fiberDtoBind(ctx *fiber.Ctx, dto interface{}) error {
	map1 := fiberGetMap(ctx)

	if len(map1) < 1 {
		map1 = map[string]interface{}{"__UnknowKey__": ""}
	}

	return mapx.BindToDto(map1, dto)
}

func fiberGetUploadedFile(ctx *fiber.Ctx, formFieldName string) *multipart.FileHeader {
	if fh, err := ctx.FormFile(formFieldName); err != nil {
		return fh
	}

	return nil
}

func fiberSendOutput(ctx *fiber.Ctx, payload ResponsePayload, err error) error {
	if err != nil {
		handler := FiberErrorHandler()
		_ = handler(ctx, err)
		return nil
	}

	LogExecuteTime(ctx)
	AddPoweredBy(ctx)

	if payload == nil {
		ctx.Type("html", "utf8")
		ctx.SendString("unsupported response payload found")
		return nil
	}

	statusCode, contents := payload.GetContents()

	if statusCode >= 400 {
		ctx.Type("html", "utf8")
		ctx.Status(500).Send([]byte{})
		return nil
	}

	if pl, ok := payload.(AttachmentResponse); ok {
		pl.AddSpecifyHeaders(ctx)
		ctx.Send(pl.Buffer())
		return nil
	}

	if pl, ok := payload.(ImageResponse); ok {
		ctx.Set(fiber.HeaderContentType, pl.GetContentType())
		ctx.Send(pl.Buffer())
		return nil
	}

	contentType := payload.GetContentType()

	if contentType != "" {
		ctx.Set(fiber.HeaderContentType, contentType)
	}

	ctx.SendString(contents)
	return nil
}

func fiberFormParam(ctx *fiber.Ctx, key string) string {
	buf := ctx.Request().PostArgs().Peek(key)

	if len(buf) > 0 {
		return string(utils.CopyBytes(buf))
	}

	mf, err := ctx.Request().MultipartForm()

	if err == nil && mf != nil {
		values := mf.Value[key]

		if len(values) > 0 {
			return values[0]
		}
	}

	return ""
}

func fiberGetMapWithRules(arg0 interface{}, rules []string) map[string]interface{} {
	var ctx *fiber.Ctx
	var srcMap map[string]interface{}

	if _ctx, ok := arg0.(*fiber.Ctx); ok && _ctx != nil {
		ctx = _ctx
	} else if map1, ok := arg0.(map[string]interface{}); ok && len(map1) > 0 {
		srcMap = map1
	}

	dstMap := map[string]interface{}{}
	re1 := regexp.MustCompile(`:[^:]+$`)
	re2 := regexp.MustCompile(`:[0-9]+$`)

	for _, s1 := range rules {
		typ := 1
		mode := 2
		dv := ""

		if strings.HasPrefix(s1, "i:") {
			s1 = strings.TrimPrefix(s1, "i:")
			typ = 2

			if re1.MatchString(s1) {
				dv = stringx.SubstringAfterLast(s1, ":")
				s1 = stringx.SubstringBeforeLast(s1, ":")
			}
		} else if strings.HasPrefix(s1, "d:") {
			s1 = strings.TrimPrefix(s1, "d:")
			typ = 3

			if re1.MatchString(s1) {
				dv = stringx.SubstringAfterLast(s1, ":")
				s1 = stringx.SubstringBeforeLast(s1, ":")
			}
		} else if strings.HasPrefix(s1, "s:") {
			s1 = strings.TrimPrefix(s1, "s:")

			if re2.MatchString(s1) {
				mode = castx.ToInt(stringx.SubstringAfterLast(s1, ":"), 2)
				s1 = stringx.SubstringBeforeLast(s1, ":")
			}
		} else if re2.MatchString(s1) {
			mode = castx.ToInt(stringx.SubstringAfterLast(s1, ":"), 2)
			s1 = stringx.SubstringBeforeLast(s1, ":")
		}

		if s1 == "" || strings.Contains(s1, ":") {
			continue
		}

		srcKey := s1
		dstKey := s1

		if strings.Contains(s1, "#") {
			srcKey = stringx.SubstringBefore(s1, "#")
			dstKey = stringx.SubstringAfter(s1, "#")
		}

		var paramValue interface{}

		if ctx != nil {
			paramValue = fiberFormParam(ctx, srcKey)

			if paramValue == "" {
				paramValue = ctx.Query(srcKey)
			}
		} else if len(srcMap) > 0 {
			paramValue = srcMap[srcKey]
		}

		switch typ {
		case 1:
			value := castx.ToString(paramValue)

			if mode != 0 {
				dstMap[dstKey] = stringx.StripTags(value)
			} else {
				dstMap[dstKey] = value
			}
		case 2:
			var value int

			if n1, err := castx.ToIntE(dv); err == nil {
				value = castx.ToInt(paramValue, n1)
			} else {
				value = castx.ToInt(paramValue)
			}

			dstMap[dstKey] = value
		case 3:
			var value float64

			if n1, err := castx.ToFloat64E(dv); err == nil {
				value = castx.ToFloat64(paramValue, n1)
			} else {
				value = castx.ToFloat64(paramValue)
			}

			dstMap[dstKey] = numberx.ToDecimalString(value)
		}
	}

	return dstMap
}
