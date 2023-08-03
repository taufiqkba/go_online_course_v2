package usecase

import (
	"go_online_course_v2/internal/user/dto"
	"go_online_course_v2/internal/user/usecase"
	"go_online_course_v2/pkg/response"
)

type RegisterUseCase interface {
	Register(dto dto.UserRequestBody) *response.Errors
}

type registerUseCase struct {
	userUseCase usecase.UserUseCase
}

func (useCase *registerUseCase) Register(dto dto.UserRequestBody) *response.Errors {
	_, err := useCase.userUseCase.Create(dto)

	if err != nil {
		return err
	}

	//	TODO send email verified using SendGrid
	return nil
}

func NewRegisterUseCase(userUseCase usecase.UserUseCase) RegisterUseCase {
	return &registerUseCase{userUseCase: userUseCase}
}
