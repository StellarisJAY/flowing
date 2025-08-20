package kb

import (
	"flowing/api/handler/common"
	"flowing/global"
	model "flowing/internal/model/kb"
	service "flowing/internal/service/kb"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateKnowledgeBase(c *gin.Context) {
	var req model.CreateKnowledgeBaseReq
	if err := c.ShouldBindJSON(&req); err != nil {
		panic(global.ErrBadRequest(err))
	}
	if err := service.CreateKnowledgeBase(c, req); err != nil {
		panic(err)
	}
	c.JSON(200, common.Ok())
}

func UpdateKnowledgeBase(c *gin.Context) {
	var req model.UpdateKnowledgeBaseReq
	if err := c.ShouldBindJSON(&req); err != nil {
		panic(global.ErrBadRequest(err))
	}
	if err := service.UpdateKnowledgeBase(c, req); err != nil {
		panic(err)
	}
	c.JSON(200, common.Ok())
}

func ListKnowledgeBase(c *gin.Context) {
	var query model.KnowledgeBaseQuery
	if err := c.ShouldBindQuery(&query); err != nil {
		panic(global.ErrBadRequest(err))
	}
	res, total, err := service.ListKnowledgeBase(c, query)
	if err != nil {
		panic(err)
	}
	c.JSON(200, common.PageResp(res, total))
}

func DeleteKnowledgeBase(c *gin.Context) {
	id, err := strconv.ParseInt(c.Query("id"), 10, 64)
	if err != nil {
		panic(global.ErrBadRequest(err))
	}
	if err := service.DeleteKnowledgeBase(c, id); err != nil {
		panic(err)
	}
	c.JSON(200, common.Ok())
}
