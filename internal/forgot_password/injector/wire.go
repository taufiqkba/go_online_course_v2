//go:build wireinject
// +build wireinject

package injector

import (
	"github.com/google/wire"
	"go_online_course_v2/internal/forgot_password/delivery/http"
	"go_online_course_v2/internal/forgot_password/repository"
	"go_online_course_v2/internal/forgot_password/usecase"
	repository2 "go_online_course_v2/internal/user/repository"
	usecase2 "go_online_course_v2/internal/user/usecase"
	"go_online_course_v2/pkg/mail/mailersend"
	"gorm.io/gorm"
)

func InitializedService(db *gorm.DB) *http.ForgotPasswordHandler {
	wire.Build(
		http.NewForgotPasswordHandler,
		repository.NewForgotPasswordRepository,
		usecase.NewForgotPasswordUseCase,
		repository2.NewUserRepository,
		usecase2.NewUserUseCase,
		mailersend.NewMailUseCase,
	)
	return &http.ForgotPasswordHandler{}
}
