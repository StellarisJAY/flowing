package monitor

import (
	"context"
	"flowing/global"
	"flowing/internal/model/common"
	"flowing/internal/model/monitor"
	"flowing/internal/repository"
	"fmt"
	"time"
)

func CreateDatasource(ctx context.Context, req monitor.CreateDatasourceReq) error {
	ds := monitor.Datasource{
		Name:        req.Name,
		Type:        req.Type,
		Username:    req.Username,
		Password:    req.Password,
		Host:        req.Host,
		Port:        req.Port,
		Database:    req.Database,
		Description: req.Description,
		Code:        req.Code,
	}
	// TODO 测试连接
	if err := monitor.CreateDatasource(ctx, ds); err != nil {
		return global.NewError(500, "创建数据源失败", err)
	}
	return nil
}

func ListDatasource(ctx context.Context, query monitor.DatasourceQuery) ([]monitor.Datasource, int64, error) {
	dss, total, err := monitor.ListDatasource(ctx, query)
	if err != nil {
		return nil, 0, global.NewError(500, "获取数据源列表失败", err)
	}
	return dss, total, nil
}

func UpdateDatasource(ctx context.Context, req monitor.UpdateDatasourceReq) error {
	ds := monitor.Datasource{
		BaseModel:   common.BaseModel{Id: req.ID},
		Name:        req.Name,
		Type:        req.Type,
		Username:    req.Username,
		Password:    req.Password,
		Host:        req.Host,
		Port:        req.Port,
		Database:    req.Database,
		Description: req.Description,
	}
	// TODO 检查数据源类型是否与使用该数据源的任务匹配
	if err := monitor.UpdateDatasource(ctx, ds); err != nil {
		return global.NewError(500, "更新数据源失败", err)
	}
	return nil
}

func DeleteDatasource(ctx context.Context, id int64) error {
	// TODO 检查数据源是否被使用
	return repository.Tx(ctx, func(c context.Context) error {
		if err := repository.DB(ctx).Delete(&monitor.Datasource{}, id).Error; err != nil {
			return global.NewError(500, "删除数据源失败", err)
		}
		return nil
	})
}

func PingDatasource(ctx context.Context, model monitor.PingDatasourceReq) (int64, error) {
	pingStart := time.Now()
	switch model.Type {
	case monitor.DatasourceTypeMysql:
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", model.Username, model.Password, model.Host, model.Port, model.Database)
		if err := repository.PingMySQL(dsn); err != nil {
			return 0, global.NewError(500, "连接数据源失败", err)
		}
	case monitor.DatasourceTypeMilvus:
	// TODO 连接Milvus
	default:
		return 0, global.NewError(500, "不支持的数据源类型", nil)
	}
	return time.Since(pingStart).Milliseconds(), nil
}
