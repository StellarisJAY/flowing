package ai

import (
	"flowing/api/handler/common"
	"flowing/global"
	model "flowing/internal/model/ai"
	service "flowing/internal/service/ai"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateProviderModel(c *gin.Context) {
	var req model.CreateProviderModelReq
	if err := c.ShouldBindJSON(&req); err != nil {
		panic(global.ErrBadRequest(err))
	}
	if err := service.CreateProviderModel(c, req); err != nil {
		panic(err)
	}
	c.JSON(200, common.Ok())
}

func UpdateProviderModel(c *gin.Context) {
	var req model.UpdateProviderModelReq
	if err := c.ShouldBindJSON(&req); err != nil {
		panic(global.ErrBadRequest(err))
	}
	if err := service.UpdateProviderModel(c, req); err != nil {
		panic(err)
	}
	c.JSON(200, common.Ok())
}

func DeleteProviderModel(c *gin.Context) {
	id := c.Query("id")
	modelId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		panic(global.ErrBadRequest(err))
	}
	if err := service.DeleteProviderModel(c, modelId); err != nil {
		panic(err)
	}
	c.JSON(200, common.Ok())
}

func ListProviderModel(c *gin.Context) {
	var req model.ProviderModelQuery
	if err := c.ShouldBindQuery(&req); err != nil {
		panic(global.ErrBadRequest(err))
	}
	result, total, err := service.ListProviderModels(c, req)
	if err != nil {
		panic(err)
	}
	c.JSON(200, common.PageResp(result, total))
}
