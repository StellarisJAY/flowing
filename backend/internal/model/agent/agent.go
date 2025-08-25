package agent

import (
	"context"
	"flowing/internal/model/common"
	"flowing/internal/repository"
	"flowing/internal/repository/db"
)

type Type string

const (
	TypeSimple   Type = "simple"
	TypeWorkflow Type = "workflow"
)

type Agent struct {
	common.BaseModel
	Name        string `json:"name" gorm:"column:name;type:varchar(16);not null;"`                // 名称
	Description string `json:"description" gorm:"column:description;type:varchar(255);not null;"` // 描述
	Type        Type   `json:"type" gorm:"column:type;type:varchar(16);not null;"`                // 类型
	Config      string `json:"config" gorm:"column:config;type:json;not null;"`                   // 配置
	Public      *bool  `json:"public" gorm:"column:public;type:boolean;not null;default:0"`       // 是否公开
	ConfigValue any    `json:"-" gorm:"-"`
}

func (Agent) TableName() string {
	return "flowing_agent"
}

type ListAgentQuery struct {
	CreateBy int64
	Name     string `json:"name" form:"name"`
	Type     Type   `json:"type" form:"type"`
	Private  bool   `json:"private" form:"private"`
	common.BaseQuery
}

type CreateAgentReq struct {
	Name        string `json:"name" form:"name" binding:"required"`
	Type        Type   `json:"type" form:"type" binding:"required"`
	Description string `json:"description" form:"description" binding:"required"`
	Public      *bool  `json:"public" form:"public" binding:"required"`
}

func CreateAgent(ctx context.Context, agent *Agent) error {
	return repository.DB(ctx).Create(agent).Error
}

func UpdateAgent(ctx context.Context, agent *Agent) error {
	return repository.DB(ctx).Model(&Agent{}).Where("id = ?", agent.Id).Updates(agent).Error
}

func ListAgent(ctx context.Context, query ListAgentQuery) ([]*Agent, int64, error) {
	var agents []*Agent
	var total int64

	d := repository.DB(ctx).Model(&Agent{}).
		Select("id, name, description, type, create_by, created_at, updated_at, public")
	if query.Private {
		d = d.Where("create_by = ?", query.CreateBy)
	} else {
		d = d.Where("create_by = ? OR public = ?", query.CreateBy, true)
	}
	if query.Name != "" {
		d = d.Where("name LIKE ?", "%"+query.Name+"%")
	}
	if query.Type != "" {
		d = d.Where("type = ?", query.Type)
	}
	if err := d.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	err := d.Scopes(db.Page(query.Page, query.PageNum, query.PageSize)).Find(&agents).Error
	if err != nil {
		return nil, 0, err
	}
	return agents, total, nil
}

func GetAgentDetail(ctx context.Context, id int64) (*Agent, error) {
	var agent *Agent
	err := repository.DB(ctx).Model(&Agent{}).Where("id = ?", id).First(&agent).Error
	if err != nil {
		return nil, err
	}
	return agent, nil
}
