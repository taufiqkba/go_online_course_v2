package repository

import (
	"go_online_course_v2/internal/discount/entity"
	"go_online_course_v2/pkg/response"
	"go_online_course_v2/pkg/utils"
	"gorm.io/gorm"
)

type DiscountRepository interface {
	FindAll(offset int, limit int) []entity.Discount
	FindByID(id int) (*entity.Discount, *response.Errors)
	FindByCode(code string) (*entity.Discount, *response.Errors)
	Create(entity entity.Discount) (*entity.Discount, *response.Errors)
	Update(entity entity.Discount) (*entity.Discount, *response.Errors)
	Delete(entity entity.Discount) *response.Errors
}

type discountRepository struct {
	db *gorm.DB
}

func (repository *discountRepository) FindAll(offset int, limit int) []entity.Discount {
	var discounts []entity.Discount
	repository.db.Scopes(utils.Paginate(offset, limit)).Find(&discounts)
	return discounts
}

func (repository *discountRepository) FindByID(id int) (*entity.Discount, *response.Errors) {
	var discount entity.Discount
	if err := repository.db.First(&discount, id).Error; err != nil {
		return nil, &response.Errors{Code: 500, Err: err}
	}
	return &discount, nil
}

func (repository *discountRepository) FindByCode(code string) (*entity.Discount, *response.Errors) {
	var discount entity.Discount

	if err := repository.db.Where("code = ?", code).First(&discount).Error; err != nil {
		return nil, &response.Errors{Code: 500, Err: err}
	}
	return &discount, nil
}

func (repository *discountRepository) Create(entity entity.Discount) (*entity.Discount, *response.Errors) {
	if err := repository.db.Create(&entity).Error; err != nil {
		return nil, &response.Errors{Code: 500, Err: err}
	}
	return &entity, nil
}

func (repository *discountRepository) Update(entity entity.Discount) (*entity.Discount, *response.Errors) {
	if err := repository.db.Save(&entity).Error; err != nil {
		return nil, &response.Errors{Code: 500, Err: err}
	}
	return &entity, nil
}

func (repository *discountRepository) Delete(entity entity.Discount) *response.Errors {
	if err := repository.db.Delete(&entity).Error; err != nil {
		return &response.Errors{Code: 500, Err: err}
	}
	return nil
}

func NewDiscountRepository(db *gorm.DB) DiscountRepository {
	return &discountRepository{db: db}
}
