package entity

import (
	entity3 "go_online_course_v2/internal/discount/entity"
	entity2 "go_online_course_v2/internal/order_detail/entity"
	"go_online_course_v2/internal/user/entity"
	"gorm.io/gorm"
	"time"
)

type Order struct {
	ID           int64                 `json:"id"`
	UserID       *int64                `json:"user_id"`
	User         *entity.User          `json:"user" gorm:"foreignKey:UserID;references:ID"`
	OrderDetails []entity2.OrderDetail `json:"order_details"`
	DiscountID   *int64                `json:"discount_id"`
	Discount     *entity3.Discount     `json:"discount" gorm:"foreignKey:DiscountID;references:ID"`
	CheckoutLink string                `json:"checkout_link"`
	ExternalID   string                `json:"external_id"`
	Price        int64                 `json:"price"`
	TotalPrice   int64                 `json:"total_price"`
	Status       string                `json:"status"`
	CreatedByID  *int64                `json:"created_by" gorm:"column:created_by"`
	CreatedBy    *entity.User          `json:"-" gorm:"foreignKey:CreatedByID;references:ID"`
	UpdatedByID  *int64                `json:"updated_by" gorm:"column:updated_by"`
	UpdatedBy    *entity.User          `json:"-" gorm:"foreignKey:UpdatedByID;references:ID"`
	CreatedAt    *time.Time            `json:"created_at"`
	UpdatedAT    *time.Time            `json:"updated_at"`
	DeletedAt    gorm.DeletedAt        `json:"deleted_at"`
}
