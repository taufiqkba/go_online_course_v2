package usecase

import (
	"go_online_course_v2/internal/product/dto"
	"go_online_course_v2/internal/product/entity"
	"go_online_course_v2/internal/product/repository"
	repository2 "go_online_course_v2/internal/product_category/repository"
	"go_online_course_v2/pkg/media/cloudinary"
	"go_online_course_v2/pkg/response"
)

type ProductUseCase interface {
	FindAll(offset int, limit int) []entity.Product
	FindByID(id int) (*entity.Product, *response.Errors)
	Create(dto dto.ProductRequestBody) (*entity.Product, *response.Errors)
	Update(id int, dto dto.ProductRequestBody) (*entity.Product, *response.Errors)
	Delete(id int) *response.Errors
	TotalCountProduct() int64
}

type productUseCase struct {
	repository                repository.ProductRepository
	productCategoryRepository repository2.ProductCategoryRepository
	media                     cloudinary.Media
}

func (useCase *productUseCase) TotalCountProduct() int64 {
	//TODO implement me
	panic("implement me")
}

func (useCase *productUseCase) FindAll(offset int, limit int) []entity.Product {
	return useCase.repository.FindAll(offset, limit)
}

func (useCase *productUseCase) FindByID(id int) (*entity.Product, *response.Errors) {
	return useCase.repository.FindByID(id)
}

func (useCase *productUseCase) Create(dto dto.ProductRequestBody) (*entity.Product, *response.Errors) {
	var product = entity.Product{
		ProductCategoryID: &dto.ProductCategoryID,
		Title:             dto.Title,
		Description:       dto.Description,
		IsHighlighted:     dto.IsHighlighted,
		Price:             int64(dto.Price),
		CreatedByID:       dto.CreatedBy,
	}

	//	check image
	if dto.Image != nil {
		image, err := useCase.media.Upload(*dto.Image)
		if err != nil {
			return nil, err
		}

		if image != nil {
			product.Image = image
		}
	}

	//	check video
	if dto.Video != nil {
		video, err := useCase.media.Upload(*dto.Video)
		if err != nil {
			return nil, err
		}

		if video != nil {
			product.Video = video
		}
	}

	data, err := useCase.repository.Create(product)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (useCase *productUseCase) Update(id int, dto dto.ProductRequestBody) (*entity.Product, *response.Errors) {
	//	find product by id
	product, err := useCase.repository.FindByID(id)
	if err != nil {
		return nil, err
	}

	product.ProductCategoryID = &dto.ProductCategoryID
	product.Title = dto.Title
	product.Description = dto.Description
	product.IsHighlighted = dto.IsHighlighted
	product.Price = int64(dto.Price)
	product.UpdatedByID = dto.UpdatedBy

	//	check image
	if dto.Image != nil {
		image, err := useCase.media.Upload(*dto.Image)
		if err != nil {
			return nil, err
		}

		// check old image
		if product.Image != nil {
			//	delete
			_, err := useCase.media.Delete(*product.Image)
			if err != nil {
				return nil, err
			}
		}

		if image != nil {
			product.Image = image
		}
	}

	//	Check video
	if dto.Video != nil {
		video, err := useCase.media.Upload(*dto.Video)
		if err != nil {
			return nil, err
		}

		//	check old video
		if product.Video != nil {
			_, err := useCase.media.Delete(*product.Video)
			if err != nil {
				return nil, err
			}
		}

		if video != nil {
			product.Video = video
		}
	}
	data, err := useCase.repository.Update(*product)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (useCase *productUseCase) Delete(id int) *response.Errors {
	//	find product by id
	product, err := useCase.repository.FindByID(id)
	if err != nil {
		return err
	}
	err = useCase.repository.Delete(*product)
	if err != nil {
		return err
	}
	return nil
}

func NewProductUseCase(repository repository.ProductRepository, productCategoryRepository repository2.ProductCategoryRepository, media cloudinary.Media) ProductUseCase {
	return &productUseCase{repository: repository, productCategoryRepository: productCategoryRepository, media: media}
}
