//go:build wireinject
// +build wireinject

package injector

import (
	"go_online_course_v2/internal/user/repository"
	userUseCase "go_online_course_v2/internal/user/usecase"
	"go_online_course_v2/internal/verification_email/delivery/http"
	"go_online_course_v2/internal/verification_email/usecase"

	"github.com/google/wire"
	"gorm.io/gorm"
)

func InitializedService(db *gorm.DB) *http.VerificationEmailHandler {
	wire.Build(
		http.NewVerificationEmailHandler,
		usecase.NewVerificationEmailUseCase,
		repository.NewUserRepository,
		userUseCase.NewUserUseCase,
	)

	return &http.VerificationEmailHandler{}
}
