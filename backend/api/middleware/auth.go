package middleware

import (
	"encoding/json"
	"flowing/global"
	"flowing/internal/model/system"
	"flowing/internal/repository"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("X-Access-Token")
		if token == "" {
			panic(global.ErrUnauthorized)
		}
		claims := new(jwt.RegisteredClaims)
		tok, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(repository.Config().Jwt.Secret), nil
		})
		if err != nil || !tok.Valid {
			panic(global.ErrUnauthorized)
		}
		result, err := repository.Redis().Get(c, claims.ID).Result()
		if err != nil || result == "" {
			panic(global.ErrUnauthorized)
		}

		userInfo := system.User{}
		if err := json.Unmarshal([]byte(result), &userInfo); err != nil {
			panic(global.ErrUnauthorized)
		}
		c.Set("user", userInfo)
		c.Next()
	}
}
