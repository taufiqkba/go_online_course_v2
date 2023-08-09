package usecase

import (
	dto2 "go_online_course_v2/internal/register/dto"
	"go_online_course_v2/internal/user/dto"
	"go_online_course_v2/internal/user/usecase"
	"go_online_course_v2/pkg/mail/mailersend"
	"go_online_course_v2/pkg/mail/sendgrid"
	"go_online_course_v2/pkg/response"
)

type RegisterUseCase interface {
	Register(dto dto.UserRequestBody) *response.Errors
}

type registerUseCase struct {
	userUseCase usecase.UserUseCase
	mail        sendgrid.Mail
	mailer      mailersend.Mail
}

func (useCase *registerUseCase) Register(dto dto.UserRequestBody) *response.Errors {
	user, err := useCase.userUseCase.Create(dto)
	if err != nil {
		return err
	}

	//	TODO send email verified using SendGrid
	data := dto2.EmailVerification{
		Subject:          "Verification Account",
		Email:            dto.Email,
		Name:             dto.Name,
		VerificationCode: user.CodeVerified,
	}

	//use goroutine
	go useCase.mailer.SendVerificationMailer(dto.Email, data)
	return nil
}

func NewRegisterUseCase(userUseCase usecase.UserUseCase, mail sendgrid.Mail, mailer mailersend.Mail) RegisterUseCase {
	return &registerUseCase{userUseCase: userUseCase, mail: mail, mailer: mailer}
}
