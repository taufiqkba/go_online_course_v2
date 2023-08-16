package usecase

import (
	"go_online_course_v2/internal/discount/dto"
	"go_online_course_v2/internal/discount/entity"
	"go_online_course_v2/internal/discount/repository"
	"go_online_course_v2/pkg/response"
)

type DiscountUseCase interface {
	FindAll(offset int, limit int) []entity.Discount
	FindByID(id int) (*entity.Discount, *response.Errors)
	FindByCode(code string) (*entity.Discount, *response.Errors)
	Create(dto dto.DiscountRequestBody) (*entity.Discount, *response.Errors)
	Update(id int, dto dto.DiscountRequestBody) (*entity.Discount, *response.Errors)
	Delete(id int) *response.Errors
	UpdateRemainingQuantity(id int, quantity int, operator string)
}

type discountUseCase struct {
	repository repository.DiscountRepository
}

func (useCase *discountUseCase) UpdateRemainingQuantity(id int, quantity int, operator string) {
	//TODO implement me
	panic("implement me")
}

func (useCase *discountUseCase) FindAll(offset int, limit int) []entity.Discount {
	return useCase.repository.FindAll(offset, limit)
}

func (useCase *discountUseCase) FindByID(id int) (*entity.Discount, *response.Errors) {
	return useCase.repository.FindByID(id)
}

func (useCase *discountUseCase) FindByCode(code string) (*entity.Discount, *response.Errors) {
	return useCase.repository.FindByCode(code)
}

func (useCase *discountUseCase) Create(dto dto.DiscountRequestBody) (*entity.Discount, *response.Errors) {
	discount := entity.Discount{
		ID:                0,
		Name:              dto.Name,
		Code:              dto.Code,
		Quantity:          dto.Quantity,
		RemainingQuantity: dto.Quantity,
		Type:              dto.Type,
		Value:             dto.Value,
		CreatedByID:       dto.CreatedBy,
	}

	if dto.StartDate != nil {
		discount.StartDate = dto.StartDate
	} else if dto.EndDate != nil {
		discount.EndDate = dto.EndDate
	}

	data, err := useCase.repository.Create(discount)
	if err != nil {
		return nil, err
	}

	return data, nil

}

func (useCase *discountUseCase) Update(id int, dto dto.DiscountRequestBody) (*entity.Discount, *response.Errors) {
	//	find discount by id
	discount, err := useCase.repository.FindByID(id)
	if err != nil {
		return nil, err
	}

	discount.Name = dto.Name
	discount.Code = dto.Code
	discount.Quantity = dto.Quantity
	discount.RemainingQuantity = dto.Quantity
	discount.Type = dto.Type
	discount.Value = dto.Value
	discount.UpdatedByID = dto.UpdatedBy

	if dto.StartDate != nil {
		discount.StartDate = dto.StartDate
	} else if dto.EndDate != nil {
		discount.EndDate = dto.EndDate
	}

	data, err := useCase.repository.Update(*discount)
	if err != nil {
		return nil, err
	}
	return data, err
}

func (useCase *discountUseCase) Delete(id int) *response.Errors {
	//	find discount by id
	discount, err := useCase.repository.FindByID(id)
	if err != nil {
		return err
	}
	err = useCase.repository.Delete(*discount)
	if err != nil {
		return err
	}
	return nil
}

func NewDiscountUseCase(repository repository.DiscountRepository) DiscountUseCase {
	return &discountUseCase{repository: repository}
}
