package api

import (
	"flowing/api/handler/sys"
	"flowing/api/handler/system"
	"flowing/api/middleware"
	"github.com/gin-gonic/gin"
	"io"
)

func InitRouter(e *gin.Engine) {
	g := e.Group("/api")
	g.Use(gin.CustomRecoveryWithWriter(io.Discard, middleware.Recovery()))
	g.Use(middleware.CORS())
	{
		g.POST("/login", sys.Login)
		g.GET("/captcha", sys.GetCaptcha)
	}
	{
		u := g.Group("/user")
		u.Use(middleware.Auth())
		u.GET("/list", system.ListUser)
		u.POST("/create", system.CreateUser)
	}
	{
		r := g.Group("/role")
		r.Use(middleware.Auth())
		r.GET("/list", system.ListRole)
		r.POST("/create", system.CreateRole)
	}
	{
		m := g.Group("/menu")
		m.Use(middleware.Auth())
		m.GET("/list", system.ListMenuTree)
		m.POST("/create", system.CreateMenu)
	}
}
