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
	{
		r := g.Group("/role")
		r.GET("/list", system.ListRole)
		r.POST("/create", system.CreateRole)
	}
	{
		m := g.Group("/menu")
		m.GET("/list", system.ListMenuTree)
		m.POST("/create", system.CreateMenu)
	}
}
