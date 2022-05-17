package gin

import (
	"github.com/gin-gonic/gin"
)

//panic
var fRecovery gin.HandlerFunc

//api request logger
var fLogger gin.HandlerFunc

func init() {

	fRecovery = gin.RecoveryWithWriter(vertical_log.GetLogger("gin-error").Fptr)

	fLogger = func(c *gin.Context) {
		vertical_log.GetLogger("api-access").Printf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}
}
