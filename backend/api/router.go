package api

import (
	"flowing/api/handler/ai"
	"flowing/api/handler/kb"
	"flowing/api/handler/monitor"
	"flowing/api/handler/sys"
	"flowing/api/handler/system"
	"flowing/api/middleware"
	"io"

	"github.com/gin-gonic/gin"
)

func InitRouter(e *gin.Engine) {
	e.Use(middleware.CORS()) // 跨域中间件
	g := e.Group("/api")
	g.Use(gin.CustomRecoveryWithWriter(io.Discard, middleware.Recovery())) // 自定义恢复中间件
	{
		g.POST("/login", sys.Login)       // 登录
		g.GET("/captcha", sys.GetCaptcha) // 获取验证码
	}
	{
		s := g.Group("/sys")
		s.Use(middleware.Auth())
		s.GET("/menus", sys.GetUserMenus)                // 获取用户菜单
		s.GET("/permissions", sys.GetUserAllPermissions) // 获取用户所有权限
	}
	{
		u := g.Group("/user")
		u.Use(middleware.Auth())
		u.GET("/list", system.ListUser)            // 获取用户列表
		u.POST("/create", system.CreateUser)       // 创建用户
		u.PUT("/update", system.UpdateUser)        // 更新用户
		u.DELETE("/delete/:id", system.DeleteUser) // 删除用户
		u.GET("/detail", system.GetUserDetail)     // 获取用户详情
	}
	{
		r := g.Group("/role")
		r.Use(middleware.Auth())
		r.GET("/list", system.ListRole)         // 获取角色列表
		r.POST("/create", system.CreateRole)    // 创建角色
		r.POST("/grant", system.CreateUserRole) // 为用户授权角色
		r.POST("/menus", system.SaveRoleMenus)  // 保存角色菜单
		r.PUT("/update", system.UpdateRole)     // 更新角色
		r.DELETE("/delete", system.DeleteRole)  // 删除角色
		r.GET("/menus", system.GetRoleMenus)    // 获取角色菜单
	}
	{
		m := g.Group("/menu")
		m.Use(middleware.Auth())
		m.GET("/list", system.ListMenuTree)    // 获取菜单列表
		m.POST("/create", system.CreateMenu)   // 创建菜单
		m.PUT("/update", system.UpdateMenu)    // 更新菜单
		m.DELETE("/delete", system.DeleteMenu) // 删除菜单
	}
	{
		p := g.Group("/ai/provider")
		p.Use(middleware.Auth())
		p.POST("/create", ai.CreateProvider) // 创建模型供应商
		p.GET("/list", ai.ListProvider)      // 获取模型供应商列表
	}
	{
		m := g.Group("/ai/model")
		m.Use(middleware.Auth())
		m.POST("/create", ai.CreateProviderModel)   // 创建模型
		m.GET("/list", ai.ListProviderModel)        // 获取模型列表
		m.DELETE("/delete", ai.DeleteProviderModel) // 删除模型
		m.PUT("/update", ai.UpdateProviderModel)    // 更新模型
	}
	{
		d := g.Group("/dict")
		d.Use(middleware.Auth())
		d.POST("/item/create", system.CreateDictItem)       // 创建字典项
		d.POST("/create", system.CreateDict)                // 创建字典
		d.GET("/list", system.ListDict)                     // 获取字典列表
		d.GET("/item/list", system.ListDictItem)            // 获取字典项列表
		d.GET("/item/list/code", system.ListDictItemByCode) // 获取字典项列表
		d.PUT("/update", system.UpdateDict)                 // 更新字典
		d.PUT("/item/update", system.UpdateDictItem)        // 更新字典项
		d.DELETE("/delete", system.DeleteDict)              // 删除字典
		d.DELETE("/item/delete", system.DeleteDictItem)     // 删除字典项
	}
	{
		m := g.Group("/monitor")
		m.Use(middleware.Auth())
		m.GET("/performance", monitor.GetSystemMetrics) // 获取系统指标

		m.GET("/datasource/list", monitor.ListDatasource)        // 获取数据源列表
		m.POST("/datasource/create", monitor.CreateDatasource)   // 创建数据源
		m.PUT("/datasource/update", monitor.UpdateDatasource)    // 更新数据源
		m.DELETE("/datasource/delete", monitor.DeleteDatasource) // 删除数据源
		m.POST("/datasource/ping", monitor.PingDatasource)       // 测试连接
	}
	{
		k := g.Group("/kb")
		k.Use(middleware.Auth())
		k.POST("/create", kb.CreateKnowledgeBase) // 创建知识库
		k.GET("/list", kb.ListKnowledgeBase)      // 获取知识库列表
		k.PUT("/update", kb.UpdateKnowledgeBase)  // 更新知识库
	}
}
