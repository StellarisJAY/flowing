package system

import (
	"flowing/api/handler/common"
	"flowing/global"
	model "flowing/internal/model/system"
	service "flowing/internal/service/system"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var user model.CreateUserReq
	if err := c.ShouldBindJSON(&user); err != nil {
		panic(global.ErrBadRequest)
	}
	if err := service.CreateUser(c, user); err != nil {
		panic(err)
	}
	c.JSON(200, common.Ok())
}

func ListUser(c *gin.Context) {
	var query model.UserQuery
	if err := c.ShouldBindQuery(&query); err != nil {
		panic(global.ErrBadRequest)
	}
	users, total, err := service.ListUser(c, query)
	if err != nil {
		panic(err)
	}
	c.JSON(200, common.PageResp(users, total))
}
