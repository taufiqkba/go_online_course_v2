package entity

import (
	"gorm.io/gorm"
	"time"
)

type Admin struct {
	ID          int64          `json:"id"`
	Name        string         `json:"name"`
	Email       string         `json:"email"`
	Password    string         `json:"password"`
	CreatedByID *int64         `json:"created_by" gorm:"column:created_by"`
	CreatedBy   *Admin         `json:"-" gorm:"foreignKey:CreatedByID;references:ID"`
	UpdatedByID *int64         `json:"updated_by" gorm:"column:updated_by"`
	UpdatedBy   *Admin         `json:"-" gorm:"foreignKey:UpdatedByID;references:ID"`
	CreatedAt   *time.Time     `json:"created_at"`
	UpdatedAt   *time.Time     `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at"`
}
