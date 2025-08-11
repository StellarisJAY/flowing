package system

import (
	"flowing/api/handler/common"
	"flowing/global"
	model "flowing/internal/model/system"
	service "flowing/internal/service/system"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateMenu(c *gin.Context) {
	var req model.CreateMenuReq
	if err := c.ShouldBindJSON(&req); err != nil {
		panic(global.ErrBadRequest(err))
	}
	if err := service.CreateMenu(c, req); err != nil {
		panic(err)
	}
	c.JSON(200, common.Ok())
}

func ListMenuTree(c *gin.Context) {
	var query model.MenuQuery
	if err := c.ShouldBindQuery(&query); err != nil {
		panic(global.ErrBadRequest(err))
	}
	menus, err := service.ListMenuTree(c, query)
	if err != nil {
		panic(err)
	}
	c.JSON(200, common.OkWithData(menus))
}

func UpdateMenu(c *gin.Context) {
	var req model.UpdateMenuReq
	if err := c.ShouldBindJSON(&req); err != nil {
		panic(global.ErrBadRequest(err))
	}
	if err := service.UpdateMenu(c, req); err != nil {
		panic(err)
	}
	c.JSON(200, common.Ok())
}

func DeleteMenu(c *gin.Context) {
	id := c.Query("id")
	menuId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		panic(global.ErrBadRequest(err))
	}
	if err := service.DeleteMenu(c, menuId); err != nil {
		panic(err)
	}
	c.JSON(200, common.Ok())
}
