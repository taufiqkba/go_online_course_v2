package usecase

import (
	"go_online_course_v2/internal/email_verification/dto"
	dto2 "go_online_course_v2/internal/user/dto"
	"go_online_course_v2/internal/user/usecase"
	"go_online_course_v2/pkg/response"
	"time"
)

type VerificationEmailUseCase interface {
	VerificationCode(dto dto.VerificationEmailRequestBody) *response.Errors
}

type verificationEmailUseCase struct {
	userUseCase usecase.UserUseCase
}

func (useCase *verificationEmailUseCase) VerificationCode(dto dto.VerificationEmailRequestBody) *response.Errors {
	user, err := useCase.userUseCase.FindOneByCodeVerified(dto.CodeVerified)
	if err != nil {
		return err
	}

	timeNow := time.Now()

	dataUpdateUser := dto2.UserUpdateRequestBody{
		EmailVerifiedAt: &timeNow,
	}
	_, err = useCase.userUseCase.Update(int(user.ID), dataUpdateUser)
	if err != nil {
		return err
	}
	return nil
}

func NewVerificationEmailUseCase(userUseCase usecase.UserUseCase) VerificationEmailUseCase {
	return &verificationEmailUseCase{userUseCase: userUseCase}
}
