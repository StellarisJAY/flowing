package common

import (
	"flowing/internal/repository"
	"time"

	"gorm.io/gorm"
)

type BaseModel struct {
	Id        int64     `json:"id,string" gorm:"column:id;primary_key;autoIncrement:false"`
	CreatedAt time.Time `json:"createdAt" gorm:"column:created_at;"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"column:updated_at;"`
}

func (m *BaseModel) BeforeCreate(tx *gorm.DB) error {
	m.Id = repository.Snowflake().Generate().Int64()
	return nil
}

func (m *BaseModel) BeforeUpdate(tx *gorm.DB) error {
	m.UpdatedAt = time.Now()
	return nil
}
