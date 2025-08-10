package ai

import (
	"context"
	"flowing/internal/model/common"
	"flowing/internal/repository"
	"flowing/internal/repository/db"
)

type ProviderType string

const (
	ProviderTypeOpenAI ProviderType = "openai" // OpenAI
	ProviderTypeOllama ProviderType = "ollama" // Ollama
	ProviderTypeTongyi ProviderType = "tongyi" // 通义
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
	return repository.DB().WithContext(ctx).Create(&model).Error
}

func ListProviders(ctx context.Context, query ProviderQuery) ([]*Provider, int64, error) {
	var providers []*Provider
	var total int64
	d := repository.DB().WithContext(ctx).Model(&Provider{})
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
