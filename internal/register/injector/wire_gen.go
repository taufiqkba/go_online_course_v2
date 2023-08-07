// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package injector

import (
	"go_online_course_v2/internal/register/delivery/http"
	usecase2 "go_online_course_v2/internal/register/usecase"
	"go_online_course_v2/internal/user/repository"
	"go_online_course_v2/internal/user/usecase"
	"gorm.io/gorm"
)

// Injectors from wire.go:

func InitializedService(db *gorm.DB) *http.RegisterHandler {
	userRepository := repository.NewUserRepository(db)
	userUseCase := usecase.NewUserUseCase(userRepository)
	registerUseCase := usecase2.NewRegisterUseCase(userUseCase)
	registerHandler := http.NewRegisterHandler(registerUseCase)
	return registerHandler
}