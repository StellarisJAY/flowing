package ai

import (
	"context"
	"encoding/json"
	"errors"
	"flowing/global"
	"flowing/internal/model/ai"
	"flowing/internal/model/ai/provider"
	"flowing/internal/repository"
	"strings"
)

func CreateProvider(ctx context.Context, req ai.CreateProviderReq) error {
	p := ai.Provider{
		ProviderName:   req.ProviderName,
		ProviderType:   req.ProviderType,
		ProviderConfig: req.ProviderConfig,
	}
	// 校验配置
	if err := ValidateProviderConfig(ai.ProviderType(req.ProviderType), req.ProviderConfig); err != nil {
		return global.NewError(400, "配置校验失败", err)
	}
	// 获取模型提供方默认配置
	providerTemplate := ai.DefaultProviderTemplate[p.ProviderType]
	return repository.Tx(ctx, func(c context.Context) error {
		if err := ai.CreateProvider(c, &p); err != nil {
			return global.NewError(500, "新增提供方失败", err)
		}
		if len(providerTemplate.Models) == 0 {
			return nil
		}
		// 为提供方增加默认模型
		providerModels := make([]ai.ProviderModel, 0, len(providerTemplate.Models))
		modelEnabled := true
		for _, model := range providerTemplate.Models {
			providerModels = append(providerModels, ai.ProviderModel{
				ProviderId: p.Id,
				ModelName:  model.Name,
				ModelType:  ai.ModelType(model.ModelType),
				Enable:     &modelEnabled,
			})
		}
		err := repository.DB(c).Model(&ai.ProviderModel{}).
			CreateInBatches(providerModels, len(providerModels)).
			Error
		if err != nil {
			return global.NewError(500, "新增模型失败", err)
		}
		return nil
	})
}

func ListProviders(ctx context.Context, query ai.ProviderQuery) ([]*ai.Provider, int64, error) {
	list, total, err := ai.ListProviders(ctx, query)
	if err != nil {
		return nil, 0, global.NewError(500, "查询提供方失败", err)
	}
	return list, total, nil
}

func ValidateProviderConfig(providerType ai.ProviderType, data string) error {
	switch providerType {
	case ai.ProviderTypeOpenAI:
		var config provider.OpenAIProviderConfig
		if err := json.Unmarshal([]byte(data), &config); err != nil {
			return err
		}
		// 校验配置
		if config.ApiKey == "" || config.BaseUrl == "" {
			return errors.New("apiKey或baseUrl不能为空")
		}
		// 校验baseUrl
		if !strings.HasPrefix(config.BaseUrl, "http") {
			return errors.New("baseUrl格式错误")
		}
		return nil
	case ai.ProviderTypeDashscope:
		config, err := provider.DashscopeFromJSON([]byte(data))
		if err != nil {
			return err
		}
		// 校验apiKey
		if config.ApiKey == "" {
			return errors.New("apiKey不能为空")
		}
		return nil
	default:
		return errors.New("不支持的提供方类型")
	}
}
