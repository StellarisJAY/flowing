package middleware

import (
	"errors"
	"flowing/api/handler/common"
	"flowing/global"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log/slog"
	"runtime/debug"
)

func Recovery() gin.RecoveryFunc {
	return func(c *gin.Context, err any) {
		if err != nil && err.(error) != nil {
			c.JSON(500, errorMessage(err.(error)))
			c.Abort()
		}
	}
}

func errorMessage(err error) common.BaseResp {
	var e global.Error
	if errors.As(err, &e) {
		slog.Error("发生错误", "error", e.Internal, "stack", string(debug.Stack()))
		return common.Resp(e.Code, e.Message)
	}
	switch {
	case errors.Is(err, gorm.ErrRecordNotFound):
		return common.Resp(404, "record not found")
	case errors.Is(err, gorm.ErrDuplicatedKey):
		return common.Resp(422, "duplicated record")
	default:
		slog.Error("发生错误", "error", err, "stack", string(debug.Stack()))
		return common.Resp(500, "internal server error")
	}
}
