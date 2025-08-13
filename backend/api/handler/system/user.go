package system

import (
	"flowing/api/handler/common"
	"flowing/global"
	model "flowing/internal/model/system"
	service "flowing/internal/service/system"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var user model.CreateUserReq
	if err := c.ShouldBindJSON(&user); err != nil {
		panic(global.ErrBadRequest(err))
	}
	if err := service.CreateUser(c, user); err != nil {
		panic(err)
	}
	c.JSON(200, common.Ok())
}

func ListUser(c *gin.Context) {
	var query model.UserQuery
	if err := c.ShouldBindQuery(&query); err != nil {
		panic(global.ErrBadRequest(err))
	}
	users, total, err := service.ListUser(c, query)
	if err != nil {
		panic(err)
	}
	c.JSON(200, common.PageResp(users, total))
}

func UpdateUser(c *gin.Context) {
	var user model.UpdateUserReq
	if err := c.ShouldBindJSON(&user); err != nil {
		panic(global.ErrBadRequest(err))
	}
	if err := service.UpdateUser(c, user); err != nil {
		panic(err)
	}
	c.JSON(200, common.Ok())
}

func DeleteUser(c *gin.Context) {
	id := c.Param("id")
	userId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		panic(global.ErrBadRequest(err))
	}
	if err := service.DeleteUser(c, userId); err != nil {
		panic(err)
	}
	c.JSON(200, common.Ok())
}

func GetUserDetail(c *gin.Context) {
	id := c.Query("id")
	userId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		panic(global.ErrBadRequest(err))
	}
	user, err := service.GetUserDetail(c, userId)
	if err != nil {
		panic(err)
	}
	c.JSON(200, common.OkWithData(user))
}
