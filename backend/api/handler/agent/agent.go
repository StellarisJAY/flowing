package agent

import (
	"flowing/api/handler/common"
	"flowing/global"
	model "flowing/internal/model/agent"
	service "flowing/internal/service/agent"

	"github.com/gin-gonic/gin"
)

func CreateAgent(c *gin.Context) {
	var req model.CreateAgentReq
	if err := c.ShouldBindJSON(&req); err != nil {
		panic(global.ErrBadRequest(err))
	}
	if err := service.CreateAgent(c, req); err != nil {
		panic(err)
	}
	c.JSON(200, common.Ok())
}

func ListAgent(c *gin.Context) {
	var req model.ListAgentQuery
	if err := c.ShouldBindQuery(&req); err != nil {
		panic(global.ErrBadRequest(err))
	}
	res, total, err := service.ListAgent(c, req)
	if err != nil {
		panic(err)
	}
	c.JSON(200, common.PageResp(res, total))
}
