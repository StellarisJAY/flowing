package ai

import (
	"context"
	"encoding/json"
	"errors"
	"flowing/global"
	"flowing/internal/model/ai"
	"flowing/internal/model/ai/provider"
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
	if err := ai.CreateProvider(ctx, p); err != nil {
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
