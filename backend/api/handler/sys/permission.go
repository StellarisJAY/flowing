package sys

import (
	"flowing/api/handler/common"
	sysmodel "flowing/internal/model/system"
	service "flowing/internal/service/system"
	"github.com/gin-gonic/gin"
)

func GetUserMenus(c *gin.Context) {
	user, _ := c.Get("user")
	userInfo := user.(sysmodel.User)
	menus, err := service.GetUserMenus(c, userInfo.Id)
	if err != nil {
		panic(err)
	}
	c.JSON(200, common.OkWithData(menus))
}

func GetUserAllPermissions(c *gin.Context) {
	user, _ := c.Get("user")
	userInfo := user.(sysmodel.User)
	permission, err := service.GetUserAllPermissions(c, userInfo.Id)
	if err != nil {
		panic(err)
	}
	c.JSON(200, common.OkWithData(permission))
}
