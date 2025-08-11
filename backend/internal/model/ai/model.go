package ai

import (
	"context"
	"flowing/internal/model/common"
	"flowing/internal/repository"
)

type ModelType string

const (
	ModelTypeLLM       ModelType = "llm"       // 大语言模型
	ModelTypeEmbedding ModelType = "embedding" // 嵌入模型
	ModelTypeReranking ModelType = "reranking" // 重排模型
)

// ProviderModel 模型提供方模型
type ProviderModel struct {
	common.BaseModel
	ProviderId  int64     `json:"providerId,string" gorm:"column:provider_id;type:bigint;not null"`
	ModelName   string    `json:"modelName" gorm:"column:model_name"`     // 模型名称
	ModelType   ModelType `json:"modelType" gorm:"column:model_type"`     // 模型类型
	ModelConfig string    `json:"modelConfig" gorm:"column:model_config"` // 模型配置
}

func (a *ProviderModel) TableName() string {
	return "ai_provider_model"
}

func CreateProviderModel(ctx context.Context, model ProviderModel) error {
	return repository.DB(ctx).Create(&model).Error
}
