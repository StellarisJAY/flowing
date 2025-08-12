package system

import (
	"flowing/api/handler/common"
	"flowing/global"
	model "flowing/internal/model/system"
	service "flowing/internal/service/system"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateRole(c *gin.Context) {
	var req model.CreateRoleReq
	if err := c.ShouldBindJSON(&req); err != nil {
		panic(global.ErrBadRequest(err))
	}
	if err := service.CreateRole(c, req); err != nil {
		panic(err)
	}
	c.JSON(200, common.Ok())
}

func ListRole(c *gin.Context) {
	var query model.RoleQuery
	if err := c.ShouldBindQuery(&query); err != nil {
		panic(global.ErrBadRequest(err))
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
		panic(global.ErrBadRequest(err))
	}
	if err := service.CreateUserRole(c, req); err != nil {
		panic(err)
	}
	c.JSON(200, common.Ok())
}

func SaveRoleMenus(c *gin.Context) {
	var req model.SaveRoleMenuReq
	if err := c.ShouldBindJSON(&req); err != nil {
		panic(global.ErrBadRequest(err))
	}
	if err := service.SaveRoleMenus(c, req); err != nil {
		panic(err)
	}
	c.JSON(200, common.Ok())
}

func UpdateRole(c *gin.Context) {
	var req model.UpdateRoleReq
	if err := c.ShouldBindJSON(&req); err != nil {
		panic(global.ErrBadRequest(err))
	}
	if err := service.UpdateRole(c, req); err != nil {
		panic(err)
	}
	c.JSON(200, common.Ok())
}

func DeleteRole(c *gin.Context) {
	id := c.Query("id")
	roleId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		panic(global.ErrBadRequest(err))
	}
	if err := service.DeleteRole(c, roleId); err != nil {
		panic(err)
	}
	c.JSON(200, common.Ok())
}

func GetRoleMenus(c *gin.Context) {
	id := c.Query("id")
	roleId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		panic(global.ErrBadRequest(err))
	}
	menus, err := service.GetRoleMenus(c, roleId)
	if err != nil {
		panic(err)
	}
	c.JSON(200, common.OkWithData(menus))
}
