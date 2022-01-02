package goboot

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lamgor666/goboot-common/AppConf"
)

func GinMidRecover() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if AppConf.GetBoolean("logging.logMiddlewareRun") {
			RuntimeLogger().Info("middleware run: mgboot.MidRecover")
		}

		defer func() {
			r := recover()

			if r == nil {
				return
			}

			var err error

			if ex, ok := r.(error); ok {
				err = ex
			} else {
				err = fmt.Errorf("%v", r)
			}

			if err == nil {
				return
			}

			ginSendOutput(ctx, nil, err)
		}()

		ctx.Next()
	}
}
