package ai

import (
	"context"
	"errors"
	"flowing/internal/model/ai/provider"
	"flowing/internal/model/common"
	"flowing/internal/repository"
	"flowing/internal/repository/db"
)

type ProviderType string

const (
	ProviderTypeOpenAI    ProviderType = "openai"    // OpenAI
	ProviderTypeOllama    ProviderType = "ollama"    // Ollama
	ProviderTypeDashscope ProviderType = "dashscope" // 通义
)

// Provider 模型提供方
type Provider struct {
	common.BaseModel
	ProviderName   string          `json:"providerName" gorm:"column:provider_name"`     // 提供方名称
	ProviderType   string          `json:"providerType" gorm:"column:provider_type"`     // 提供方类型
	ProviderConfig string          `json:"providerConfig" gorm:"column:provider_config"` // 提供方配置
	ProviderModels []ProviderModel `json:"providerModels" gorm:"foreignKey:ProviderId;references:Id"`
}

func (m *Provider) TableName() string {
	return "ai_provider"
}

type ProviderQuery struct {
	common.BaseQuery
	ProviderType ProviderType `json:"providerType" form:"providerType"`
	ProviderName string       `json:"providerName" form:"providerName"`
}

type CreateProviderReq struct {
	ProviderName   string `json:"providerName" binding:"required"`
	ProviderType   string `json:"providerType" binding:"required"`
	ProviderConfig string `json:"providerConfig" binding:"required"`
}

func CreateProvider(ctx context.Context, model Provider) error {
	return repository.DB(ctx).Create(&model).Error
}

func ListProviders(ctx context.Context, query ProviderQuery) ([]*Provider, int64, error) {
	var providers []*Provider
	var total int64
	d := repository.DB(ctx).Model(&Provider{})
	if query.ProviderType != "" {
		d = d.Where("provider_type = ?", query.ProviderType)
	}
	if query.ProviderName != "" {
		d = d.Where("provider_name LIKE %?%", query.ProviderName)
	}
	if err := d.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	err := d.Scopes(db.Page(query.Page, query.PageNum, query.PageSize)).Preload("ProviderModels").Find(&providers).Error
	if err != nil {
		return nil, 0, err
	}
	return providers, total, nil
}

func GetProviderConfig(providerType ProviderType, config string) (any, error) {
	switch providerType {
	case ProviderTypeOpenAI:
		cfg, err := provider.OpenAIFromJSON([]byte(config))
		if err != nil {
			return nil, err
		}
		return cfg, nil
	case ProviderTypeDashscope:
		cfg, err := provider.DashscopeFromJSON([]byte(config))
		if err != nil {
			return nil, err
		}
		return cfg, nil
	default:
		return nil, errors.New("不支持的提供方类型")
	}
}
