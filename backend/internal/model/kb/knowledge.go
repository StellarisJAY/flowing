package kb

import (
	"context"
	"flowing/internal/model/common"
	"flowing/internal/repository"
	"flowing/internal/repository/db"
)

type KnowledgeBase struct {
	common.BaseModel
	Name           string `json:"name" gorm:"type:varchar(255);not null"`            // 知识库名称
	Description    string `json:"description" gorm:"type:varchar(255);not null"`     // 介绍
	DatasourceId   int64  `json:"datasourceId,string" gorm:"type:bigint;not null"`   // 数据源ID
	EmbeddingModel int64  `json:"embeddingModel,string" gorm:"type:bigint;not null"` // 嵌入模型ID
	Enable         *bool  `json:"enable" gorm:"type:tinyint(1);not null;default:0"`

	DatasourceName     string `json:"datasourceName" gorm:"-"`
	DatasourceType     string `json:"datasourceType" gorm:"-"`
	EmbeddingModelName string `json:"embeddingModelName" gorm:"-"`
}

func (k *KnowledgeBase) TableName() string {
	return "ai_knowledge_base"
}

type CreateKnowledgeBaseReq struct {
	Name           string `json:"name" binding:"required"`
	Description    string `json:"description" binding:"required"`
	DatasourceId   int64  `json:"datasourceId,string" binding:"required"`
	EmbeddingModel int64  `json:"embeddingModel,string" binding:"required"`
	Enable         *bool  `json:"enable" binding:"required"`
}

type UpdateKnowledgeBaseReq struct {
	Id          int64  `json:"id,string" binding:"required"`
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
	Enable      *bool  `json:"enable" binding:"required"`
}

type KnowledgeBaseQuery struct {
	common.BaseQuery
	Name string `json:"name" form:"name"`
}

func CreateKnowledgeBase(ctx context.Context, model KnowledgeBase) error {
	return repository.DB(ctx).Create(&model).Error
}

func UpdateKnowledgeBase(ctx context.Context, model KnowledgeBase) error {
	return repository.DB(ctx).Model(&model).Where("id = ?", model.Id).Updates(model).Error
}

func DeleteKnowledgeBase(ctx context.Context, id int64) error {
	return repository.DB(ctx).Delete(&KnowledgeBase{}, "id = ?", id).Error
}

func GetKnowledgeBase(ctx context.Context, id int64) (*KnowledgeBase, error) {
	var model *KnowledgeBase
	if err := repository.DB(ctx).Model(&KnowledgeBase{}).Where("id = ?", id).First(&model).Error; err != nil {
		return nil, err
	}
	return model, nil
}

func ListKnowledgeBase(ctx context.Context, query KnowledgeBaseQuery) ([]*KnowledgeBase, int64, error) {
	var list []*KnowledgeBase
	var total int64
	d := repository.DB(ctx).Model(&KnowledgeBase{}).
		InnerJoins("JOIN monitor_datasource md ON md.id = datasource_id").
		InnerJoins("JOIN ai_provider_model apm ON apm.id = embedding_model").
		Select("ai_knowledge_base.*, md.name as datasource_name, apm.model_name as embedding_model_name")
	if query.Name != "" {
		d = d.Where("ai_knowledge_base.name LIKE ?", "%"+query.Name+"%")
	}
	if err := d.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	if err := d.Scopes(db.Page(query.Page, query.PageNum, query.PageSize)).Find(&list).Error; err != nil {
		return nil, 0, err
	}
	return list, total, nil
}
