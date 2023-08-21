package entity

import (
	entity2 "go_online_course_v2/internal/product/entity"
	"go_online_course_v2/internal/user/entity"
	"gorm.io/gorm"
	"time"
)

type ClassRoom struct {
	ID          int64            `json:"id"`
	UserID      int64            `json:"user_id"`
	User        *entity.User     `json:"user" gorm:"foreignKey:UserID;references:ID"`
	ProductID   *int64           `json:"product_id"`
	Product     *entity2.Product `json:"product" gorm:"foreignKey:ProductID;references:ID"`
	CreatedByID *int64           `json:"created_by" gorm:"column:created_by"`
	CreatedBy   *entity.User     `json:"-" gorm:"foreignKey:CreatedByID;references:ID"`
	UpdatedByID *int64           `json:"updated_by" gorm:"column:updated_by"`
	UpdatedBy   *entity.User     `json:"-" gorm:"foreignKey:UpdatedByID;references:ID"`
	CreatedAt   *time.Time       `json:"created_at"`
	UpdatedAt   *time.Time       `json:"updated_at"`
	DeletedAt   gorm.DeletedAt   `json:"deleted_at"`
}
