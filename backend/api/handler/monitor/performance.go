package monitor

import (
	"flowing/api/handler/common"
	"flowing/global"

	metricsmodel "flowing/internal/model/monitor"

	"github.com/gin-gonic/gin"
)

func GetSystemMetrics(c *gin.Context) {
	metrics, err := metricsmodel.GatherSystemMetrics()
	if err != nil {
		panic(global.NewError(500, "获取系统指标失败", err))
	}
	c.JSON(200, common.OkWithData(metrics))
}
