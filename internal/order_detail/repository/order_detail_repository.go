package repository

import (
	"go_online_course_v2/internal/order_detail/entity"
	"go_online_course_v2/pkg/response"
	"gorm.io/gorm"
)

type OrderDetailRepository interface {
	Create(entity entity.OrderDetail) (*entity.OrderDetail, *response.Errors)
}

type orderDetailRepository struct {
	db *gorm.DB
}

func (repository *orderDetailRepository) Create(entity entity.OrderDetail) (*entity.OrderDetail, *response.Errors) {
	if err := repository.db.Create(&entity).Error; err != nil {
		return nil, &response.Errors{
			Code: 500,
			Err:  err,
		}
	}
	return &entity, nil
}

func NewOrderDetailRepository(db *gorm.DB) OrderDetailRepository {
	return &orderDetailRepository{db: db}
}
