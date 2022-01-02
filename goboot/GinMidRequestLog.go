package goboot

import (
	"github.com/gin-gonic/gin"
	"github.com/lamgor666/goboot-common/AppConf"
	"strings"
	"time"
)

func GinMidRequestLog() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if AppConf.GetBoolean("logging.logMiddlewareRun") {
			RuntimeLogger().Info("middleware run: mgboot.MidRequestLog")
		}

		ctx.Set("ExecStart", time.Now())

		if !RequestLogEnabled() {
			ctx.Next()
			return
		}

		logger := RequestLogLogger()
		sb := strings.Builder{}
		sb.WriteString(ginGetMethod(ctx))
		sb.WriteString(" ")
		sb.WriteString(ginGetRequestUrl(ctx, true))
		sb.WriteString(" from ")
		sb.WriteString(ginGetClientIp(ctx))
		logger.Info(sb.String())

		if LogRequestBody() {
			rawBody := ginGetRawBody(ctx)

			if len(rawBody) > 0 {
				logger.Debugf(string(rawBody))
			}
		}

		ctx.Next()
	}
}
