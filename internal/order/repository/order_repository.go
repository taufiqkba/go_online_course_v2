package repository

import (
	"go_online_course_v2/internal/order/entity"
	"go_online_course_v2/pkg/response"
	"go_online_course_v2/pkg/utils"
	"gorm.io/gorm"
)

type OrderRepository interface {
	FindAllByUserID(userID int, offset int, limit int) []entity.Order
	FindOneByExternalID(externalID string) (*entity.Order, *response.Errors)
	FindOneByID(id int) (*entity.Order, *response.Errors)
	Create(entity entity.Order) (*entity.Order, *response.Errors)
	Update(entity entity.Order) (*entity.Order, *response.Errors)
	TotalCountOrder() int64
}

type orderRepository struct {
	db *gorm.DB
}

func (repository *orderRepository) FindAllByUserID(userID int, offset int, limit int) []entity.Order {
	var orders []entity.Order

	repository.db.Scopes(utils.Paginate(offset, limit)).
		Preload("OrderDetails.Product").
		Where("user_id = ?", userID).
		Find(&orders)
	return orders
}

func (repository *orderRepository) FindOneByExternalID(externalID string) (*entity.Order, *response.Errors) {
	var order entity.Order

	if err := repository.db.
		Preload("OrderDetails.Product").
		Where("external_id = ?", externalID).
		First(&order).Error; err != nil {
		return nil, &response.Errors{
			Code: 500,
			Err:  err,
		}
	}
	return &order, nil
}

func (repository *orderRepository) FindOneByID(id int) (*entity.Order, *response.Errors) {
	var order entity.Order

	if err := repository.db.
		Preload("OrderDetails.Product").
		First(&order, id).Error; err != nil {
		return nil, &response.Errors{
			Code: 500,
			Err:  err,
		}
	}
	return &order, nil
}

func (repository *orderRepository) Create(entity entity.Order) (*entity.Order, *response.Errors) {
	if err := repository.db.Create(&entity).Error; err != nil {
		return nil, &response.Errors{
			Code: 500,
			Err:  err,
		}
	}
	return &entity, nil
}

func (repository *orderRepository) Update(entity entity.Order) (*entity.Order, *response.Errors) {
	if err := repository.db.Save(&entity).Error; err != nil {
		return nil, &response.Errors{
			Code: 500,
			Err:  err,
		}
	}
	return &entity, nil
}

func (repository *orderRepository) TotalCountOrder() int64 {
	var order entity.Order
	var totalOrder int64

	repository.db.Model(&order).Count(&totalOrder)
	return totalOrder
}

func NewOrderRepository(db *gorm.DB) OrderRepository {
	return &orderRepository{db: db}
}
