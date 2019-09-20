package middleware

import (
	"github.com/ECEHive/myHive-backend/model"
	"github.com/ECEHive/myHive-backend/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

var error_handler_logger = util.GetLogger("error_handler")

func ErrorHandler(ctx *gin.Context) {
	var logId = util.SecureRandomString(10)
	ctx.Set("LogId", logId)
	ctx.Header("SPEC-Log-Id", logId)
	ctx.Writer.Header().Set("Content-Type", "application/json; charset=utf-8")
	ctx.Next()
	if err, exists := ctx.Get("error"); exists {
		v, ok := err.(model.ErrorResponse)
		if !ok {
			error_handler_logger.Error("Failed to Cast error to ErrorResponse")
			ctx.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		ctx.AbortWithStatusJSON(v.StatusCode, v)
		error_handler_logger.WithField("SPEC-Log-Id", logId).Errorf("%s\n%+v\n%+v", ctx.Request.URL.Path, *ctx.Request, v)
	}
}
