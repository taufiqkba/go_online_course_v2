package usecase

import (
	"errors"
	"go_online_course_v2/internal/cart/dto"
	"go_online_course_v2/internal/cart/entity"
	"go_online_course_v2/internal/cart/repository"
	"go_online_course_v2/pkg/response"
)

type CartUseCase interface {
	FindAllByUserID(userID int, offset int, limit int) []entity.Cart
	FindByID(id int) (*entity.Cart, *response.Errors)
	Create(dto dto.CartRequestBody) (*entity.Cart, *response.Errors)
	Update(id int, dto dto.CartUpdateRequestBody) (*entity.Cart, *response.Errors)
	Delete(id int, userID int) *response.Errors
	DeleteByUserID(userID int) *response.Errors
}

type cartUseCase struct {
	repository repository.CartRepository
}

func (useCase *cartUseCase) FindAllByUserID(userID int, offset int, limit int) []entity.Cart {
	return useCase.repository.FindAllByUserID(userID, offset, limit)
}

func (useCase *cartUseCase) FindByID(id int) (*entity.Cart, *response.Errors) {
	return useCase.repository.FindOneByID(id)
}

func (useCase *cartUseCase) Create(dto dto.CartRequestBody) (*entity.Cart, *response.Errors) {
	cart := entity.Cart{
		UserID:      &dto.UserID,
		ProductID:   &dto.ProductID,
		Quantity:    1,
		IsChecked:   true,
		CreatedByID: &dto.CreatedBy,
	}

	data, err := useCase.repository.Create(cart)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (useCase *cartUseCase) Update(id int, dto dto.CartUpdateRequestBody) (*entity.Cart, *response.Errors) {
	//	find by user id
	cart, err := useCase.repository.FindOneByID(id)
	if err != nil {
		return nil, err
	}
	//	validate
	if *cart.UserID != *dto.UserID {
		return nil, &response.Errors{
			Code: 400,
			Err:  errors.New("this cart is not yours"),
		}
	}
	cart.IsChecked = dto.IsChecked
	cart.UpdatedByID = dto.UserID

	updateCart, err := useCase.repository.Update(*cart)
	if err != nil {
		return nil, err
	}
	return updateCart, nil
}

func (useCase *cartUseCase) Delete(id int, userID int) *response.Errors {
	//	find data by id
	cart, err := useCase.repository.FindOneByID(id)
	if err != nil {
		return err
	}

	//check to make sure the cart is owner same with userID
	user := int64(userID)
	if *cart.UserID != user {
		return &response.Errors{
			Code: 400,
			Err:  errors.New("this is not your carts"),
		}
	}

	err = useCase.repository.Delete(*cart)
	if err != nil {
		return err
	}
	return nil
}

func (useCase *cartUseCase) DeleteByUserID(userID int) *response.Errors {
	err := useCase.repository.DeleteByUserID(userID)
	if err != nil {
		return err
	}
	return nil
}

func NewCartUseCase(repository repository.CartRepository) CartUseCase {
	return &cartUseCase{repository: repository}
}
