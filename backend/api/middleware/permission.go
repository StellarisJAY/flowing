package middleware

import (
	"flowing/global"
	sysmodel "flowing/internal/model/system"

	"github.com/gin-gonic/gin"
)

func Permission(code string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从上下文获取用户信息
		user, exists := c.Get("user")
		if !exists {
			panic(global.ErrUnauthorized)
		}
		// 转换为用户模型
		userInfo, ok := user.(sysmodel.User)
		if !ok {
			panic(global.ErrUnauthorized)
		}
		// 检查用户权限
		if !userInfo.UserPermission.HasPermission(code) {
			panic(global.ErrForbidden)
		}
		c.Next()
	}
}
