package repository

import (
	"go_online_course_v2/internal/cart/entity"
	"go_online_course_v2/pkg/response"
	"go_online_course_v2/pkg/utils"
	"gorm.io/gorm"
)

type CartRepository interface {
	FindAllByUserID(userID int, offset int, limit int) []entity.Cart
	FindOneByID(id int) (*entity.Cart, *response.Errors)
	Create(entity entity.Cart) (*entity.Cart, *response.Errors)
	Update(entity entity.Cart) (*entity.Cart, *response.Errors)
	Delete(entity entity.Cart) *response.Errors
	DeleteByUserID(id int) *response.Errors
}

type cartRepository struct {
	db *gorm.DB
}

func (repository *cartRepository) FindAllByUserID(userID int, offset int, limit int) []entity.Cart {
	var carts []entity.Cart

	repository.db.Scopes(utils.Paginate(offset, limit)).
		Preload("User").Preload("Product").
		Where("user_id = ?", userID).
		Find(&carts)
	return carts
}

func (repository *cartRepository) FindOneByID(id int) (*entity.Cart, *response.Errors) {
	var cart entity.Cart

	if err := repository.db.Preload("User").Preload("Product").Find(&cart, id).Error; err != nil {
		return nil, &response.Errors{
			Code: 500,
			Err:  err,
		}
	}
	return &cart, nil
}

func (repository *cartRepository) Create(entity entity.Cart) (*entity.Cart, *response.Errors) {
	if err := repository.db.Create(&entity).Error; err != nil {
		return nil, &response.Errors{
			Code: 500,
			Err:  err,
		}
	}
	return &entity, nil
}

func (repository *cartRepository) Update(entity entity.Cart) (*entity.Cart, *response.Errors) {
	if err := repository.db.Save(&entity).Error; err != nil {
		return nil, &response.Errors{
			Code: 500,
			Err:  err,
		}
	}
	return &entity, nil
}

func (repository *cartRepository) Delete(entity entity.Cart) *response.Errors {
	if err := repository.db.Delete(&entity).Error; err != nil {
		return &response.Errors{
			Code: 500,
			Err:  err,
		}
	}
	return nil
}

func (repository *cartRepository) DeleteByUserID(id int) *response.Errors {
	var cart entity.Cart

	if err := repository.db.Where("user_id = ?").Delete(&cart).Error; err != nil {
		return &response.Errors{
			Code: 500,
			Err:  err,
		}
	}
	return nil
}

func NewCartRepository(db *gorm.DB) CartRepository {
	return &cartRepository{db: db}
}
