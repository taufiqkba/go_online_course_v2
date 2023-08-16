package repository

import (
	"go_online_course_v2/internal/product/entity"
	"go_online_course_v2/pkg/response"
	"go_online_course_v2/pkg/utils"
	"gorm.io/gorm"
)

type ProductRepository interface {
	FindAll(offset int, limit int) []entity.Product
	FindByID(id int) (*entity.Product, *response.Errors)
	Create(entity entity.Product) (*entity.Product, *response.Errors)
	Update(entity entity.Product) (*entity.Product, *response.Errors)
	Delete(entity entity.Product) *response.Errors
	TotalCountProduct() int64
}

type productRepository struct {
	db *gorm.DB
}

func (repository *productRepository) TotalCountProduct() int64 {
	//implement me
	panic("implement me")
}

func (repository *productRepository) FindAll(offset int, limit int) []entity.Product {
	var products []entity.Product

	repository.db.Scopes(utils.Paginate(offset, limit)).Preload("ProductCategory").Find(&products)
	return products
}

func (repository *productRepository) FindByID(id int) (*entity.Product, *response.Errors) {
	var product entity.Product

	if err := repository.db.Preload("ProductCategory").First(&product, id).Error; err != nil {
		return nil, &response.Errors{
			Code: 500,
			Err:  err,
		}
	}
	return &product, nil
}

func (repository *productRepository) Create(entity entity.Product) (*entity.Product, *response.Errors) {
	if err := repository.db.Create(&entity).Error; err != nil {
		return nil, &response.Errors{
			Code: 500,
			Err:  err,
		}
	}
	return &entity, nil
}

func (repository *productRepository) Update(entity entity.Product) (*entity.Product, *response.Errors) {
	if err := repository.db.Save(&entity).Error; err != nil {
		return nil, &response.Errors{
			Code: 500,
			Err:  err,
		}
	}
	return &entity, nil
}

func (repository *productRepository) Delete(entity entity.Product) *response.Errors {
	if err := repository.db.Delete(&entity).Error; err != nil {
		return &response.Errors{
			Code: 500,
			Err:  err,
		}
	}
	return nil
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productRepository{db: db}
}
