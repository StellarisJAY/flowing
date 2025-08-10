package system

import (
	"flowing/api/handler/common"
	"flowing/global"
	model "flowing/internal/model/system"
	service "flowing/internal/service/system"
	"github.com/gin-gonic/gin"
)

func CreateDict(c *gin.Context) {
	var req model.CreateDictReq
	if err := c.ShouldBindJSON(&req); err != nil {
		panic(global.ErrBadRequest(err))
	}
	if err := service.CreateDict(c, req); err != nil {
		panic(err)
	}
	c.JSON(200, common.Ok())
}

func CreateDictItem(c *gin.Context) {
	var req model.CreateDictItemReq
	if err := c.ShouldBindJSON(&req); err != nil {
		panic(global.ErrBadRequest(err))
	}
	if err := service.CreateDictItem(c, req); err != nil {
		panic(err)
	}
	c.JSON(200, common.Ok())
}

func ListDict(c *gin.Context) {
	var req model.DictQuery
	if err := c.ShouldBindQuery(&req); err != nil {
		panic(global.ErrBadRequest(err))
	}
	list, total, err := service.ListDict(c, req)
	if err != nil {
		panic(err)
	}
	c.JSON(200, common.PageResp(list, total))
}

func ListDictItem(c *gin.Context) {
	var req model.DictItemQuery
	if err := c.ShouldBindQuery(&req); err != nil {
		panic(global.ErrBadRequest(err))
	}
	list, total, err := service.ListDictItems(c, req)
	if err != nil {
		panic(err)
	}
	c.JSON(200, common.PageResp(list, total))
}

func ListDictItemByCode(c *gin.Context) {
	var req model.ListDictItemByCodeReq
	if err := c.ShouldBindQuery(&req); err != nil {
		panic(global.ErrBadRequest(err))
	}
	list, err := service.ListDictItemByCode(c, req.Code)
	if err != nil {
		panic(err)
	}
	c.JSON(200, common.PageResp(list, int64(len(list))))
}
