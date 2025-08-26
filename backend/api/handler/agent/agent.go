package agent

import (
	"flowing/api/handler/common"
	"flowing/global"
	model "flowing/internal/model/agent"
	service "flowing/internal/service/agent"
	"strconv"

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

func UpdateConfig(c *gin.Context) {
	var req model.UpdateAgentConfigReq
	if err := c.ShouldBindJSON(&req); err != nil {
		panic(global.ErrBadRequest(err))
	}
	if err := service.UpdateConfig(c, req); err != nil {
		panic(err)
	}
	c.JSON(200, common.Ok())
}

func GetDetail(c *gin.Context) {
	id, err := strconv.ParseInt(c.Query("id"), 10, 64)
	if err != nil {
		panic(global.ErrBadRequest(err))
	}
	agent, err := service.GetAgentDetail(c, id)
	if err != nil {
		panic(err)
	}
	c.JSON(200, common.OkWithData(agent))
}
