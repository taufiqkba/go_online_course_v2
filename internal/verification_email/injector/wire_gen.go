// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package injector

import (
	"go_online_course_v2/internal/user/repository"
	"go_online_course_v2/internal/user/usecase"
	"go_online_course_v2/internal/verification_email/delivery/http"
	usecase2 "go_online_course_v2/internal/verification_email/usecase"
	"gorm.io/gorm"
)

// Injectors from wire.go:

func InitializedService(db *gorm.DB) *http.VerificationEmailHandler {
	userRepository := repository.NewUserRepository(db)
	userUseCase := usecase.NewUserUseCase(userRepository)
	verificationEmailUseCase := usecase2.NewVerificationEmailUseCase(userUseCase)
	verificationEmailHandler := http.NewVerificationEmailHandler(verificationEmailUseCase)
	return verificationEmailHandler
}
