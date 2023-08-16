package repository

import (
	"go_online_course_v2/internal/product_category/entity"
	"go_online_course_v2/pkg/response"
	"go_online_course_v2/pkg/utils"
	"gorm.io/gorm"
)

type ProductCategoryRepository interface {
	FindAll(offset int, limit int) []entity.ProductCategory
	FindByID(id int) (*entity.ProductCategory, *response.Errors)
	Create(entity entity.ProductCategory) (*entity.ProductCategory, *response.Errors)
	Update(entity entity.ProductCategory) (*entity.ProductCategory, *response.Errors)
	Delete(entity entity.ProductCategory) *response.Errors
}
type productCategoryRepository struct {
	db *gorm.DB
}

func (repository *productCategoryRepository) FindAll(offset int, limit int) []entity.ProductCategory {
	var productCategories []entity.ProductCategory

	repository.db.Scopes(utils.Paginate(offset, limit)).Find(&productCategories)
	return productCategories
}

func (repository *productCategoryRepository) FindByID(id int) (*entity.ProductCategory, *response.Errors) {
	var productCategory entity.ProductCategory

	if err := repository.db.First(&productCategory, id).Error; err != nil {
		return nil, &response.Errors{
			Code: 500,
			Err:  err,
		}
	}
	return &productCategory, nil
}

func (repository *productCategoryRepository) Create(entity entity.ProductCategory) (*entity.ProductCategory, *response.Errors) {
	if err := repository.db.Create(&entity).Error; err != nil {
		return nil, &response.Errors{
			Code: 500,
			Err:  err,
		}
	}
	return &entity, nil
}

func (repository *productCategoryRepository) Update(entity entity.ProductCategory) (*entity.ProductCategory, *response.Errors) {
	if err := repository.db.Save(&entity); err != nil {
		return nil, &response.Errors{
			Code: 500,
			Err:  nil,
		}
	}
	return &entity, nil
}

func (repository *productCategoryRepository) Delete(entity entity.ProductCategory) *response.Errors {
	if err := repository.db.Delete(&entity).Error; err != nil {
		return &response.Errors{
			Code: 500,
			Err:  err,
		}
	}
	return nil
}

func NewProductCategoryRepository(db *gorm.DB) ProductCategoryRepository {
	return &productCategoryRepository{db: db}
}
