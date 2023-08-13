//go:build wireinject
// +build wireinject

package injector

import (
	"github.com/google/wire"
	"go_online_course_v2/internal/admin/delivery/http"
	"go_online_course_v2/internal/admin/repository"
	"go_online_course_v2/internal/admin/usecase"
	"gorm.io/gorm"
)

func InitializedService(db *gorm.DB) *http.AdminHandler {

	wire.Build(
		http.NewAdminHandler,
		repository.NewAdminRepository,
		usecase.NewAdminUseCase,
	)
	return &http.AdminHandler{}
}
