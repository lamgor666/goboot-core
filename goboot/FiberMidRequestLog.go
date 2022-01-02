package goboot

import (
	"github.com/gofiber/fiber/v2"
	"github.com/lamgor666/goboot-common/AppConf"
	"strings"
	"time"
)

func FiberMidRequestLog() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		if AppConf.GetBoolean("logging.logMiddlewareRun") {
			RuntimeLogger().Info("middleware run: mgboot.MidRequestLog")
		}

		ctx.Locals("ExecStart", time.Now())

		if !RequestLogEnabled() {
			return ctx.Next()
		}

		logger := RequestLogLogger()
		sb := strings.Builder{}
		sb.WriteString(ctx.Method())
		sb.WriteString(" ")
		sb.WriteString(fiberGetRequestUrl(ctx, true))
		sb.WriteString(" from ")
		sb.WriteString(fiberGetClientIp(ctx))
		logger.Info(sb.String())

		if LogRequestBody() {
			rawBody := fiberGetRawBody(ctx)

			if len(rawBody) > 0 {
				logger.Debugf(string(rawBody))
			}
		}

		return ctx.Next()
	}
}
