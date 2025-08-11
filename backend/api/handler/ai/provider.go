package ai

import (
	"flowing/api/handler/common"
	"flowing/global"
	model "flowing/internal/model/ai"
	service "flowing/internal/service/ai"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateProvider(c *gin.Context) {
	var req model.CreateProviderReq
	if err := c.ShouldBindJSON(&req); err != nil {
		panic(global.ErrBadRequest(err))
	}
	if err := service.CreateProvider(c, req); err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, common.Ok())
}

func ListProvider(c *gin.Context) {
	var req model.ProviderQuery
	if err := c.ShouldBindQuery(&req); err != nil {
		panic(global.ErrBadRequest(err))
	}
	providers, total, err := service.ListProviders(c, req)
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, common.PageResp(providers, total))
}
