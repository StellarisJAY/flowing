package ai

import (
	"context"
	"flowing/global"
	"flowing/internal/model/ai"
)

func CreateProvider(ctx context.Context, req ai.CreateProviderReq) error {
	provider := ai.Provider{
		ProviderName:   req.ProviderName,
		ProviderType:   req.ProviderType,
		ProviderConfig: req.ProviderConfig,
	}
	// TODO 验证配置json
	if err := ai.CreateProvider(ctx, provider); err != nil {
		return global.NewError(500, "新增提供方失败", err)
	}
	return nil
}

func ListProviders(ctx context.Context, query ai.ProviderQuery) ([]*ai.Provider, int64, error) {
	list, total, err := ai.ListProviders(ctx, query)
	if err != nil {
		return nil, 0, global.NewError(500, "查询提供方失败", err)
	}
	return list, total, nil
}
