package ai

import (
	"context"
	"flowing/internal/model/common"
	"flowing/internal/repository"
	"flowing/internal/repository/db"
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
	ProviderId int64     `json:"providerId,string" gorm:"column:provider_id;type:bigint;not null"`
	ModelName  string    `json:"modelName" gorm:"column:model_name"`        // 模型名称
	ModelType  ModelType `json:"modelType" gorm:"column:model_type"`        // 模型类型
	Enable     *bool     `json:"enable" gorm:"column:enable;default:false"` // 是否启用
}

func (a *ProviderModel) TableName() string {
	return "ai_provider_model"
}

type ProviderModelQuery struct {
	common.BaseQuery
	ProviderId int64     `json:"providerId,string" form:"providerId"`
	ModelType  ModelType `json:"modelType" form:"modelType"`
	ModelName  string    `json:"modelName" form:"modelName"`
	Enable     *bool     `json:"enable" form:"enable"`
}

type CreateProviderModelReq struct {
	ProviderId int64     `json:"providerId,string" binding:"required"`
	ModelName  string    `json:"modelName" binding:"required"`
	ModelType  ModelType `json:"modelType" binding:"required"`
	Enable     *bool     `json:"enable" binding:"required"`
}

type ProviderModelDetail struct {
	common.BaseModel
	ProviderId     int64        `json:"providerId,string" gorm:"column:provider_id;type:bigint;not null"`
	ModelName      string       `json:"modelName" gorm:"column:model_name"`           // 模型名称
	ModelType      ModelType    `json:"modelType" gorm:"column:model_type"`           // 模型类型
	Enable         *bool        `json:"enable" gorm:"column:enable;default:false"`    // 是否启用
	ProviderName   string       `json:"providerName" gorm:"column:provider_name"`     // 模型提供方名称
	ProviderType   ProviderType `json:"providerType" gorm:"column:provider_type"`     // 模型提供方类型
	ProviderConfig string       `json:"providerConfig" gorm:"column:provider_config"` // 模型提供方配置
}

type ProviderModelListVo struct {
	common.BaseModel
	ProviderId   int64        `json:"providerId,string" gorm:"column:provider_id"`
	ModelName    string       `json:"modelName" gorm:"column:model_name"`        // 模型名称
	ModelType    ModelType    `json:"modelType" gorm:"column:model_type"`        // 模型类型
	Enable       *bool        `json:"enable" gorm:"column:enable;default:false"` // 是否启用
	ProviderName string       `json:"providerName" gorm:"column:provider_name"`  // 模型提供方名称
	ProviderType ProviderType `json:"providerType" gorm:"column:provider_type"`  // 模型提供方类型
}

func CreateProviderModel(ctx context.Context, model ProviderModel) error {
	return repository.DB(ctx).Create(&model).Error
}

func UpdateProviderModel(ctx context.Context, model ProviderModel) error {
	return repository.DB(ctx).Model(&model).Updates(&model).Error
}

type UpdateProviderModelReq struct {
	Id     int64 `json:"id,string" binding:"required"`
	Enable *bool `json:"enable" binding:"required"`
}

func GetProviderModel(ctx context.Context, id int64) (*ProviderModel, error) {
	var model *ProviderModel
	err := repository.DB(ctx).First(&model, id).Error
	return model, err
}

func ListProviderModel(ctx context.Context, query ProviderModelQuery) ([]*ProviderModelListVo, int64, error) {
	var models []*ProviderModelListVo
	var total int64
	d := repository.DB(ctx).Table("ai_provider_model apm").
		Joins("JOIN ai_provider ap ON ap.id = apm.provider_id").
		Select("apm.*, ap.provider_name as provider_name")
	if query.ProviderId != 0 {
		d = d.Where("provider_id = ?", query.ProviderId)
	}
	if query.ModelType != "" {
		d = d.Where("model_type = ?", query.ModelType)
	}
	if query.ModelName != "" {
		d = d.Where("model_name LIKE ?", "%"+query.ModelName+"%")
	}
	if query.Enable != nil {
		d = d.Where("enable = ?", *query.Enable)
	}
	if err := d.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	err := d.Scopes(db.Page(query.Page, query.PageNum, query.PageSize)).Find(&models).Error
	return models, total, err
}
