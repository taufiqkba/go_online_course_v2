package usecase

import (
	"errors"
	"go_online_course_v2/internal/forgot_password/dto"
	"go_online_course_v2/internal/forgot_password/entity"
	"go_online_course_v2/internal/forgot_password/repository"
	dto2 "go_online_course_v2/internal/user/dto"
	"go_online_course_v2/internal/user/usecase"
	"go_online_course_v2/pkg/mail/mailersend"
	"go_online_course_v2/pkg/response"
	"go_online_course_v2/pkg/utils"
	"time"
)

type ForgotPasswordUseCase interface {
	Create(dtoForgotPassword dto.ForgotPasswordRequestBody) (*entity.ForgotPassword, *response.Errors)
	Update(dto dto.ForgotPasswordUpdateRequestBody) (*entity.ForgotPassword, *response.Errors)
}

type forgotPasswordUseCase struct {
	repository  repository.ForgotPasswordRepository
	userUseCase usecase.UserUseCase
	mail        mailersend.Mail
}

func (useCase *forgotPasswordUseCase) Create(dtoForgotPassword dto.ForgotPasswordRequestBody) (*entity.ForgotPassword, *response.Errors) {
	user, err := useCase.userUseCase.FindByEmail(dtoForgotPassword.Email)
	if err != nil {
		return nil, err
	}

	//for security reason, response set to 200 success meanwhile is it not found
	if user == nil {
		return nil, &response.Errors{
			Code: 200,
			Err:  errors.New("success, please check your email"),
		}
	}

	dateTime := time.Now().Add(24 * 1 * time.Hour)
	forgotPassword := entity.ForgotPassword{
		UserId:    &user.ID,
		Valid:     true,
		Code:      utils.RandNumber(8),
		ExpiredAt: &dateTime,
	}

	dataForgotPassword, err := useCase.repository.Create(forgotPassword)

	//	send email forgot password verification code
	dataEmailForgotPassword := dto.ForgotPasswordEmailRequestBody{
		Subject: "Verification Code Forgot Password",
		Name:    user.Name,
		Email:   user.Email,
		Code:    forgotPassword.Code,
	}

	go useCase.mail.SendForgotPasswordMailer(user.Email, dataEmailForgotPassword)
	if err != nil {
		return nil, err
	}
	return dataForgotPassword, nil
}

func (useCase *forgotPasswordUseCase) Update(dto dto.ForgotPasswordUpdateRequestBody) (*entity.ForgotPassword, *response.Errors) {
	//	check code verification password
	code, err := useCase.repository.FindByCode(dto.Code)
	if err != nil || !code.Valid {
		return nil, &response.Errors{
			Code: 400,
			Err:  errors.New("code is invalid"),
		}
	}

	//	search user
	user, err := useCase.userUseCase.FindOneByID(int(*code.UserId))
	if err != nil {
		return nil, err
	}

	//	update password to table user
	dataUser := dto2.UserUpdateRequestBody{
		Password: &dto.Password,
	}

	_, err = useCase.userUseCase.Update(int(user.ID), dataUser)
	if err != nil {
		return nil, err
	}
	code.Valid = false

	useCase.repository.Update(*code)
	return code, nil
}

func NewForgotPasswordUseCase(repository repository.ForgotPasswordRepository, userUseCase usecase.UserUseCase, mail mailersend.Mail) ForgotPasswordUseCase {
	return &forgotPasswordUseCase{repository: repository, userUseCase: userUseCase, mail: mail}
}
