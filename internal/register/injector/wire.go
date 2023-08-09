//go:build wireinject
// +build wireinject

package injector

import (
	"github.com/google/wire"
	"go_online_course_v2/internal/register/delivery/http"
	"go_online_course_v2/internal/register/usecase"
	"go_online_course_v2/internal/user/repository"
	usecase2 "go_online_course_v2/internal/user/usecase"
	"go_online_course_v2/pkg/mail/mailersend"
	"go_online_course_v2/pkg/mail/sendgrid"
	"gorm.io/gorm"
)

func InitializedService(db *gorm.DB) *http.RegisterHandler {
	wire.Build(
		usecase.NewRegisterUseCase,
		http.NewRegisterHandler,
		repository.NewUserRepository,
		usecase2.NewUserUseCase,
		sendgrid.NewMailUseCase,
		mailersend.NewMailUseCase,
	)
	return &http.RegisterHandler{}
}
