package common

import (
	"gorm.io/gorm"
	"time"
)

type BaseModel struct {
	Id        int            `json:"id" gorm:"column:id;"`
	CreatedAt time.Time      `json:"createdAt" gorm:"column:created_at;"`
	UpdatedAt time.Time      `json:"updatedAt" gorm:"column:updated_at;"`
	DeletedAt gorm.DeletedAt `json:"deletedAt" gorm:"column:deleted_at;"`
}

func (m *BaseModel) BeforeCreate(tx *gorm.DB) error {
	return nil
}
