package middleware

import (
	"flowing/global"
	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("X-Access-Token")
		if token == "" {
			panic(global.ErrUnauthorized)
		}
		// TODO: check token, get user info
		c.Next()
	}
}
