//go:build wireinject
// +build wireinject

package injector

import (
	"github.com/google/wire"
	"go_online_course_v2/internal/user/delivery/http"
	"go_online_course_v2/internal/user/repository"
	"go_online_course_v2/internal/user/usecase"
	"gorm.io/gorm"
)

func InitializedService(db *gorm.DB) *http.UserHandler {
	wire.Build(
		repository.NewUserRepository,
		usecase.NewUserUseCase,
		http.NewUserHandler,
	)
	return &http.UserHandler{}
}
