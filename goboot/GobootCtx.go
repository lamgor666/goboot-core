package goboot

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/gofiber/fiber/v2"
	"github.com/lamgor666/goboot-common/util/castx"
	"math"
	"mime/multipart"
)

func GetMethod(arg0 interface{}) string {
	if ctx, ok := arg0.(*fiber.Ctx); ok {
		return fiberGetMethod(ctx)
	}

	if ctx, ok := arg0.(*gin.Context); ok {
		return ginGetMethod(ctx)
	}

	return ""
}

func GetHeader(arg0 interface{}, name string) string {
	if ctx, ok := arg0.(*fiber.Ctx); ok {
		return fiberGetHeader(ctx, name)
	}

	if ctx, ok := arg0.(*gin.Context); ok {
		return ginGetHeader(ctx, name)
	}

	return ""
}

func GetHeaders(arg0 interface{}) map[string]string {
	if ctx, ok := arg0.(*fiber.Ctx); ok {
		return fiberGetHeaders(ctx)
	}

	if ctx, ok := arg0.(*gin.Context); ok {
		return ginGetHeaders(ctx)
	}

	return map[string]string{}
}

func GetQueryParams(arg0 interface{}) map[string]string {
	if ctx, ok := arg0.(*fiber.Ctx); ok {
		return fiberGetQueryParams(ctx)
	}

	if ctx, ok := arg0.(*gin.Context); ok {
		return ginGetQueryParams(ctx)
	}

	return map[string]string{}
}

func GetQueryString(arg0 interface{}, urlencode ...bool) string {
	if ctx, ok := arg0.(*fiber.Ctx); ok {
		return fiberGetQueryString(ctx, urlencode...)
	}

	if ctx, ok := arg0.(*gin.Context); ok {
		return ginGetQueryString(ctx, urlencode...)
	}

	return ""
}

func GetRequestUrl(arg0 interface{}, withQueryString ...bool) string {
	if ctx, ok := arg0.(*fiber.Ctx); ok {
		return fiberGetRequestUrl(ctx, withQueryString...)
	}

	if ctx, ok := arg0.(*gin.Context); ok {
		return ginGetRequestUrl(ctx, withQueryString...)
	}

	return ""
}

func GetFormData(arg0 interface{}) map[string]string {
	if ctx, ok := arg0.(*fiber.Ctx); ok {
		return fiberGetFormData(ctx)
	}

	if ctx, ok := arg0.(*gin.Context); ok {
		return ginGetFormData(ctx)
	}

	return map[string]string{}
}

func GetClientIp(arg0 interface{}) string {
	if ctx, ok := arg0.(*fiber.Ctx); ok {
		return fiberGetClientIp(ctx)
	}

	if ctx, ok := arg0.(*gin.Context); ok {
		return ginGetClientIp(ctx)
	}

	return ""
}

func Pathvariable(arg0 interface{}, name string, defaultValue ...interface{}) string {
	var dv string

	if len(defaultValue) > 0 {
		if s1, err := castx.ToStringE(defaultValue[0]); err == nil {
			dv = s1
		}
	}

	if ctx, ok := arg0.(*fiber.Ctx); ok {
		return fiberPathvariable(ctx, name, dv)
	}

	if ctx, ok := arg0.(*gin.Context); ok {
		return ginPathvariable(ctx, name, dv)
	}

	return dv
}

func PathvariableBool(arg0 interface{}, name string, defaultValue ...interface{}) bool {
	var dv bool

	if len(defaultValue) > 0 {
		if b1, err := castx.ToBoolE(defaultValue[0]); err == nil {
			dv = b1
		}
	}

	if ctx, ok := arg0.(*fiber.Ctx); ok {
		return fiberPathvariableBool(ctx, name, dv)
	}

	if ctx, ok := arg0.(*gin.Context); ok {
		return ginPathvariableBool(ctx, name, dv)
	}

	return dv
}

func PathvariableInt(arg0 interface{}, name string, defaultValue ...interface{}) int {
	dv := math.MinInt32

	if len(defaultValue) > 0 {
		if n1, err := castx.ToIntE(defaultValue[0]); err == nil {
			dv = n1
		}
	}

	if ctx, ok := arg0.(*fiber.Ctx); ok {
		return fiberPathvariableInt(ctx, name, dv)
	}

	if ctx, ok := arg0.(*gin.Context); ok {
		return ginPathvariableInt(ctx, name, dv)
	}

	return dv
}

func PathvariableInt64(arg0 interface{}, name string, defaultValue ...interface{}) int64 {
	dv := int64(math.MinInt64)

	if len(defaultValue) > 0 {
		if n1, err := castx.ToInt64E(defaultValue[0]); err == nil {
			dv = n1
		}
	}

	if ctx, ok := arg0.(*fiber.Ctx); ok {
		return fiberPathvariableInt64(ctx, name, dv)
	}

	if ctx, ok := arg0.(*gin.Context); ok {
		return ginPathvariableInt64(ctx, name, dv)
	}

	return dv
}

func PathvariableFloat32(arg0 interface{}, name string, defaultValue ...interface{}) float32 {
	dv := float32(math.SmallestNonzeroFloat32)

	if len(defaultValue) > 0 {
		if n1, err := castx.ToFloat32E(defaultValue[0]); err == nil {
			dv = n1
		}
	}

	if ctx, ok := arg0.(*fiber.Ctx); ok {
		return fiberPathvariableFloat32(ctx, name, dv)
	}

	if ctx, ok := arg0.(*gin.Context); ok {
		return ginPathvariableFloat32(ctx, name, dv)
	}

	return dv
}

func PathvariableFloat64(arg0 interface{}, name string, defaultValue ...interface{}) float64 {
	dv := math.SmallestNonzeroFloat64

	if len(defaultValue) > 0 {
		if n1, err := castx.ToFloat64E(defaultValue[0]); err == nil {
			dv = n1
		}
	}

	if ctx, ok := arg0.(*fiber.Ctx); ok {
		return fiberPathvariableFloat64(ctx, name, dv)
	}

	if ctx, ok := arg0.(*gin.Context); ok {
		return ginPathvariableFloat64(ctx, name, dv)
	}

	return dv
}

func ReqParam(arg0 interface{}, name string, mode int, defaultValue ...interface{}) string {
	var dv string

	if len(defaultValue) > 0 {
		if s1, err := castx.ToStringE(defaultValue[0]); err == nil {
			dv = s1
		}
	}

	if ctx, ok := arg0.(*fiber.Ctx); ok {
		return fiberReqParam(ctx, name, mode, dv)
	}

	if ctx, ok := arg0.(*gin.Context); ok {
		return ginReqParam(ctx, name, mode, dv)
	}

	return dv
}

func ReqParamBool(arg0 interface{}, name string, defaultValue ...interface{}) bool {
	var dv bool

	if len(defaultValue) > 0 {
		if b1, err := castx.ToBoolE(defaultValue[0]); err == nil {
			dv = b1
		}
	}

	if ctx, ok := arg0.(*fiber.Ctx); ok {
		return fiberReqParamBool(ctx, name, dv)
	}

	if ctx, ok := arg0.(*gin.Context); ok {
		return ginReqParamBool(ctx, name, dv)
	}

	return dv
}

func ReqParamInt(arg0 interface{}, name string, defaultValue ...interface{}) int {
	dv := math.MinInt32

	if len(defaultValue) > 0 {
		if n1, err := castx.ToIntE(defaultValue[0]); err == nil {
			dv = n1
		}
	}

	if ctx, ok := arg0.(*fiber.Ctx); ok {
		return fiberReqParamInt(ctx, name, dv)
	}

	if ctx, ok := arg0.(*gin.Context); ok {
		return ginReqParamInt(ctx, name, dv)
	}

	return dv
}

func ReqParamInt64(arg0 interface{}, name string, defaultValue ...interface{}) int64 {
	dv := int64(math.MinInt64)

	if len(defaultValue) > 0 {
		if n1, err := castx.ToInt64E(defaultValue[0]); err == nil {
			dv = n1
		}
	}

	if ctx, ok := arg0.(*fiber.Ctx); ok {
		return fiberReqParamInt64(ctx, name, dv)
	}

	if ctx, ok := arg0.(*gin.Context); ok {
		return ginReqParamInt64(ctx, name, dv)
	}

	return dv
}

func ReqParamFloat32(arg0 interface{}, name string, defaultValue ...interface{}) float32 {
	dv := float32(math.SmallestNonzeroFloat32)

	if len(defaultValue) > 0 {
		if n1, err := castx.ToFloat32E(defaultValue[0]); err == nil {
			dv = n1
		}
	}

	if ctx, ok := arg0.(*fiber.Ctx); ok {
		return fiberReqParamFloat32(ctx, name, dv)
	}

	if ctx, ok := arg0.(*gin.Context); ok {
		return ginReqParamFloat32(ctx, name, dv)
	}

	return dv
}

func ReqParamFloat64(arg0 interface{}, name string, defaultValue ...interface{}) float64 {
	dv := math.SmallestNonzeroFloat64

	if len(defaultValue) > 0 {
		if n1, err := castx.ToFloat64E(defaultValue[0]); err == nil {
			dv = n1
		}
	}

	if ctx, ok := arg0.(*fiber.Ctx); ok {
		return fiberReqParamFloat64(ctx, name, dv)
	}

	if ctx, ok := arg0.(*gin.Context); ok {
		return ginReqParamFloat64(ctx, name, dv)
	}

	return dv
}

func GetJwt(arg0 interface{}) *jwt.Token {
	if ctx, ok := arg0.(*fiber.Ctx); ok {
		return fiberGetJwt(ctx)
	}

	if ctx, ok := arg0.(*gin.Context); ok {
		return ginGetJwt(ctx)
	}

	return nil
}

func GetRawBody(arg0 interface{}) []byte {
	if ctx, ok := arg0.(*fiber.Ctx); ok {
		return fiberGetRawBody(ctx)
	}

	if ctx, ok := arg0.(*gin.Context); ok {
		return ginGetRawBody(ctx)
	}

	return make([]byte, 0)
}

func GetMap(arg0 interface{}, rules ...interface{}) map[string]interface{} {
	if ctx, ok := arg0.(*fiber.Ctx); ok {
		return fiberGetMap(ctx, rules...)
	}

	if ctx, ok := arg0.(*gin.Context); ok {
		return ginGetMap(ctx, rules...)
	}

	return map[string]interface{}{}
}

func DtoBind(arg0 interface{}, dto interface{}) error {
	if ctx, ok := arg0.(*fiber.Ctx); ok {
		return fiberDtoBind(ctx, dto)
	}

	if ctx, ok := arg0.(*gin.Context); ok {
		return ginDtoBind(ctx, dto)
	}

	return nil
}

func GetUploadedFile(arg0 interface{}, formFieldName string) *multipart.FileHeader {
	if ctx, ok := arg0.(*fiber.Ctx); ok {
		return fiberGetUploadedFile(ctx, formFieldName)
	}

	if ctx, ok := arg0.(*gin.Context); ok {
		return ginGetUploadedFile(ctx, formFieldName)
	}

	return nil
}

func SendOutput(arg0 interface{}, payload ResponsePayload, err error) error {
	if ctx, ok := arg0.(*fiber.Ctx); ok {
		return fiberSendOutput(ctx, payload, err)
	}

	if ctx, ok := arg0.(*gin.Context); ok {
		ginSendOutput(ctx, payload, err)
		return nil
	}

	return nil
}
