package common

import (
	"flowing/internal/repository"
	"gorm.io/gorm"
	"time"
)

type BaseModel struct {
	Id        int64          `json:"id" gorm:"column:id;primary_key;"`
	CreatedAt time.Time      `json:"createdAt" gorm:"column:created_at;"`
	UpdatedAt time.Time      `json:"updatedAt" gorm:"column:updated_at;"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"column:deleted_at;"`
}

func (m *BaseModel) BeforeCreate(tx *gorm.DB) error {
	m.Id = repository.Snowflake().Generate().Int64()
	return nil
}
