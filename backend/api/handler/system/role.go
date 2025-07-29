package system

import (
	"flowing/api/handler/common"
	model "flowing/internal/model/system"
	service "flowing/internal/service/system"
	"github.com/gin-gonic/gin"
)

func CreateRole(c *gin.Context) {
	var req model.CreateRoleReq
	if err := c.ShouldBindJSON(&req); err != nil {
		panic(common.ErrBadRequest)
	}
	if err := service.CreateRole(c, req); err != nil {
		panic(err)
	}
	c.JSON(200, common.Ok())
}

func ListRole(c *gin.Context) {
	var query model.RoleQuery
	if err := c.ShouldBindQuery(&query); err != nil {
		panic(common.ErrBadRequest)
	}
	roles, total, err := service.ListRole(c, query)
	if err != nil {
		panic(err)
	}
	c.JSON(200, common.PageResp(roles, total))
}
