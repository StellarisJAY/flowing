package monitor

import (
	"flowing/api/handler/common"
	"flowing/global"
	monitorModel "flowing/internal/model/monitor"
	monitorService "flowing/internal/service/monitor"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateDatasource(c *gin.Context) {
	var req monitorModel.CreateDatasourceReq
	if err := c.ShouldBindJSON(&req); err != nil {
		panic(global.ErrBadRequest(err))
	}
	if err := monitorService.CreateDatasource(c, req); err != nil {
		panic(err)
	}
	c.JSON(200, common.Ok())
}

func ListDatasource(c *gin.Context) {
	var req monitorModel.DatasourceQuery
	if err := c.ShouldBindQuery(&req); err != nil {
		panic(global.ErrBadRequest(err))
	}
	list, total, err := monitorService.ListDatasource(c, req)
	if err != nil {
		panic(err)
	}
	c.JSON(200, common.PageResp(list, total))
}

func DeleteDatasource(c *gin.Context) {
	id := c.Query("id")
	dsId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		panic(global.ErrBadRequest(err))
	}
	if err := monitorService.DeleteDatasource(c, dsId); err != nil {
		panic(err)
	}
	c.JSON(200, common.Ok())
}

func UpdateDatasource(c *gin.Context) {
	var req monitorModel.UpdateDatasourceReq
	if err := c.ShouldBindJSON(&req); err != nil {
		panic(global.ErrBadRequest(err))
	}
	if err := monitorService.UpdateDatasource(c, req); err != nil {
		panic(err)
	}
	c.JSON(200, common.Ok())
}

func PingDatasource(c *gin.Context) {
	var req monitorModel.PingDatasourceReq
	if err := c.ShouldBindJSON(&req); err != nil {
		panic(global.ErrBadRequest(err))
	}
	latency, err := monitorService.PingDatasource(c, req)
	if err != nil {
		panic(err)
	}
	c.JSON(200, common.OkWithData(latency))
}
