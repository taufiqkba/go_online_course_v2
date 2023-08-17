package entity

import (
	"go_online_course_v2/internal/product/entity"
	entity2 "go_online_course_v2/internal/user/entity"
	"gorm.io/gorm"
	"time"
)

type OrderDetail struct {
	ID          int64           `json:"id"`
	OrderID     int64           `json:"order_id"`
	ProductID   int64           `json:"product_id"`
	Product     *entity.Product `json:"product" gorm:"foreignKey:ProductID;references:ID"`
	Price       int64           `json:"price"`
	CreatedByID *int64          `json:"created_by" gorm:"column:created_by"`
	CreatedBy   *entity2.User   `json:"-" gorm:"foreignKey:CreatedByID;references:ID"`
	UpdatedByID *int64          `json:"updated_by" gorm:"column:updated_by"`
	UpdatedBy   *entity2.User   `json:"-" gorm:"foreignKey:UpdatedByID;references:ID"`
	CreatedAt   *time.Time      `json:"created_at"`
	UpdatedAT   *time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt  `json:"deleted_at"`
}
