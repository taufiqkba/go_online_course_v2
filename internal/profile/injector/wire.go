//go:build wireinject
// +build wireinject

package injector

import (
	"github.com/google/wire"
	"go_online_course_v2/internal/admin/repository"
	"go_online_course_v2/internal/admin/usecase"
	repository2 "go_online_course_v2/internal/oauth/repository"
	usecase2 "go_online_course_v2/internal/oauth/usecase"
	"go_online_course_v2/internal/profile/delivery/http"
	usecase4 "go_online_course_v2/internal/profile/usecase"
	repository3 "go_online_course_v2/internal/user/repository"
	usecase3 "go_online_course_v2/internal/user/usecase"
	"gorm.io/gorm"
)

func InitializedService(db *gorm.DB) *http.ProfileHandler {
	wire.Build(
		repository.NewAdminRepository,
		usecase.NewAdminUseCase,
		repository2.NewOauthAccessTokenRepository,
		repository2.NewOauthClientRepository,
		repository2.NewOauthRefreshTokenRepository,
		usecase2.NewOauthUseCase,
		repository3.NewUserRepository,
		usecase3.NewUserUseCase,
		usecase4.NewProfileUseCase,
		http.NewProfileHandler,
	)
	return &http.ProfileHandler{}
}
