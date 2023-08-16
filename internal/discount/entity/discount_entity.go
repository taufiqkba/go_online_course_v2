package entity

import (
	"go_online_course_v2/internal/admin/entity"
	"gorm.io/gorm"
	"time"
)

type Discount struct {
	ID                int64          `json:"id"`
	Name              string         `json:"name"`
	Code              string         `json:"code"`
	Quantity          int64          `json:"quantity"`
	RemainingQuantity int64          `json:"remaining_quantity"`
	Type              string         `json:"type "`
	Value             int64          `json:"value "`
	StartDate         *time.Time     `json:"start_date"`
	EndDate           *time.Time     `json:"end_date"`
	CreatedByID       *int64         `json:"created_by" gorm:"column:created_by"`
	CreatedBy         *entity.Admin  `json:"-" gorm:"foreignKey:CreatedByID;references:ID"`
	UpdatedByID       *int64         `json:"updated_by" gorm:"column:updated_by"`
	UpdatedBy         *entity.Admin  `json:"-" gorm:"foreignKey:UpdatedByID;references:ID"`
	CreatedAt         *time.Time     `json:"created_at"`
	UpdatedAt         *time.Time     `json:"updated_at"`
	DeletedAt         gorm.DeletedAt `json:"deleted_at"`
}
