package system

import (
	"flowing/api/handler/common"
	"flowing/global"
	model "flowing/internal/model/system"
	service "flowing/internal/service/system"
	"github.com/gin-gonic/gin"
	"log/slog"
)

func CreateRole(c *gin.Context) {
	var req model.CreateRoleReq
	if err := c.ShouldBindJSON(&req); err != nil {
		panic(global.ErrBadRequest)
	}
	if err := service.CreateRole(c, req); err != nil {
		panic(err)
	}
	c.JSON(200, common.Ok())
}

func ListRole(c *gin.Context) {
	var query model.RoleQuery
	if err := c.ShouldBindQuery(&query); err != nil {
		panic(global.ErrBadRequest)
	}
	roles, total, err := service.ListRole(c, query)
	if err != nil {
		panic(err)
	}
	c.JSON(200, common.PageResp(roles, total))
}

func CreateUserRole(c *gin.Context) {
	var req model.CreateUserRoleReq
	if err := c.ShouldBindJSON(&req); err != nil {
		slog.Error("error", "e", err)
		panic(global.ErrBadRequest)
	}
	if err := service.CreateUserRole(c, req); err != nil {
		panic(err)
	}
	c.JSON(200, common.Ok())
}

func CreateRoleMenu(c *gin.Context) {
	var req model.CreateRoleMenuReq
	if err := c.ShouldBindJSON(&req); err != nil {
		panic(global.ErrBadRequest)
	}
	if err := service.CreateRoleMenu(c, req); err != nil {
		panic(err)
	}
	c.JSON(200, common.Ok())
}
