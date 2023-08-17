package usecase

import (
	"fmt"
	"go_online_course_v2/internal/order_detail/entity"
	"go_online_course_v2/internal/order_detail/repository"
)

type OrderDetailUseCase interface {
	Create(entity entity.OrderDetail)
}

type orderDetailUseCase struct {
	repository repository.OrderDetailRepository
}

func (useCase *orderDetailUseCase) Create(entity entity.OrderDetail) {
	_, err := useCase.repository.Create(entity)
	if err != nil {
		fmt.Println(err)
	}
}

func NewOrderDetailUseCase(repository repository.OrderDetailRepository) OrderDetailUseCase {
	return &orderDetailUseCase{repository: repository}
}
