package goboot

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/gofiber/fiber/v2"
	"io/ioutil"
	"strings"
)

func GinMidRequestBody() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		method := GinGetMethod(ctx)
		isPost := method == "POST"
		isPut := method == "PUT"
		isPatch := method == "PATCH"
		isDelete := method == "DELETE"
		contentType := strings.ToLower(GinGetHeader(ctx, fiber.HeaderContentType))
		isJson := (isPost || isPut || isPatch || isDelete) && strings.Contains(contentType, fiber.MIMEApplicationJSON)
		isXml := (isPost || isPut || isPatch || isDelete) && (strings.Contains(contentType, fiber.MIMEApplicationXML) || strings.Contains(contentType, fiber.MIMETextXML))

		if isJson || isXml {
			if buf, err := ctx.GetRawData(); err == nil && len(buf) > 0 {
				ctx.Set("requestRawBody", buf)
				ctx.Request.Body = ioutil.NopCloser(bytes.NewReader(buf))
			}

			return
		}

		ctx.Next()
	}
}
