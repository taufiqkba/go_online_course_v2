package usecase

import (
	"go_online_course_v2/internal/product_category/dto"
	"go_online_course_v2/internal/product_category/entity"
	"go_online_course_v2/internal/product_category/repository"
	"go_online_course_v2/pkg/media/cloudinary"
	"go_online_course_v2/pkg/response"
)

type ProductCategoryUseCase interface {
	FindAll(offset int, limit int) []entity.ProductCategory
	FindByID(id int) (*entity.ProductCategory, *response.Errors)
	Create(dto dto.ProductCategoryRequestBody) (*entity.ProductCategory, *response.Errors)
	Update(id int, dto dto.ProductCategoryRequestBody) (*entity.ProductCategory, *response.Errors)
	Delete(id int) *response.Errors
}

type productCategoryUseCase struct {
	repository repository.ProductCategoryRepository
	media      cloudinary.Media
}

func (useCase *productCategoryUseCase) FindAll(offset int, limit int) []entity.ProductCategory {
	return useCase.FindAll(offset, limit)
}

func (useCase *productCategoryUseCase) FindByID(id int) (*entity.ProductCategory, *response.Errors) {
	return useCase.FindByID(id)
}

func (useCase *productCategoryUseCase) Create(dto dto.ProductCategoryRequestBody) (*entity.ProductCategory, *response.Errors) {
	productCategory := entity.ProductCategory{
		Name:        dto.Name,
		CreatedByID: dto.CreatedBy,
	}

	if dto.Image != nil {
		//	Upload image to cloudinary
		image, err := useCase.media.Upload(*dto.Image)
		if err != nil {
			return nil, err
		}

		if image != nil {
			productCategory.Image = image
		}
	}

	data, err := useCase.repository.Create(productCategory)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (useCase *productCategoryUseCase) Update(id int, dto dto.ProductCategoryRequestBody) (*entity.ProductCategory, *response.Errors) {
	//	find by id
	productCategory, err := useCase.repository.FindByID(id)
	if err != nil {
		return nil, err
	}

	productCategory.Name = dto.Name
	productCategory.UpdatedByID = dto.UpdatedBy

	//set upload new image updated
	if dto.Image != nil {
		image, err := useCase.media.Upload(*dto.Image)

		if err != nil {
			return nil, err
		}

		//	check old image
		if productCategory.Image != nil {
			//	delete old image
			_, err := useCase.media.Delete(*productCategory.Image)
			if err != nil {
				return nil, err
			}
		}

		//	if image already
		if image != nil {
			productCategory.Image = image
		}
	}
	//	update to database
	data, err := useCase.repository.Update(*productCategory)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (useCase *productCategoryUseCase) Delete(id int) *response.Errors {
	//	Find By ID
	productCategory, err := useCase.FindByID(id)
	if err != nil {
		return err
	}

	if err := useCase.repository.Delete(*productCategory); err != nil {
		return err
	}
	return nil
}

func NewProductCategoryUseCase(repository repository.ProductCategoryRepository, media cloudinary.Media) ProductCategoryUseCase {
	return &productCategoryUseCase{repository: repository, media: media}
}
