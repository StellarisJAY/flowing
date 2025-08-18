package ai

import (
	"context"
	"flowing/global"
	model "flowing/internal/model/ai"
	"flowing/internal/model/common"
	"flowing/internal/repository"
)

func CreateProviderModel(ctx context.Context, req model.CreateProviderModelReq) error {
	pm := model.ProviderModel{
		BaseModel:  common.BaseModel{},
		ProviderId: req.ProviderId,
		ModelName:  req.ModelName,
		ModelType:  req.ModelType,
		Enable:     req.Enable,
	}
	if err := model.CreateProviderModel(ctx, pm); err != nil {
		return global.NewError(500, "创建模型失败", err)
	}
	return nil
}

func UpdateProviderModel(ctx context.Context, req model.UpdateProviderModelReq) error {
	pm := model.ProviderModel{
		BaseModel: common.BaseModel{
			Id: req.Id,
		},
		Enable: req.Enable,
	}
	if err := model.UpdateProviderModel(ctx, pm); err != nil {
		return global.NewError(500, "更新模型失败", err)
	}
	return nil
}

func ListProviderModels(ctx context.Context, query model.ProviderModelQuery) ([]*model.ProviderModel, int64, error) {
	result, total, err := model.ListProviderModel(ctx, query)
	if err != nil {
		return nil, 0, global.NewError(500, "查询模型列表失败", err)
	}
	return result, total, nil
}

func DeleteProviderModel(ctx context.Context, id int64) error {
	// TODO 检查模型是否被使用
	if err := repository.DB(ctx).Delete(&model.ProviderModel{}, "id=?", id).Error; err != nil {
		return global.NewError(500, "删除模型失败", err)
	}
	return nil
}
