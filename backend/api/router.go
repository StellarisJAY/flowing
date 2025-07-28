package api

import (
	"flowing/api/handler/system"
	"github.com/gin-gonic/gin"
)

func InitRouter(e *gin.Engine) {
	g := e.Group("/api")
	{
		u := g.Group("/user")
		u.GET("/list", system.ListUser)
		u.POST("/create", system.CreateUser)
	}
}
