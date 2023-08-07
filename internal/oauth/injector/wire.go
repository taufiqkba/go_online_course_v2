//go:build wireinject
// +build wireinject

package injector

import (
	"github.com/google/wire"
	"go_online_course_v2/internal/oauth/delivery/http"
	"go_online_course_v2/internal/oauth/repository"
	"go_online_course_v2/internal/oauth/usecase"
	repository2 "go_online_course_v2/internal/user/repository"
	usecase2 "go_online_course_v2/internal/user/usecase"
	"gorm.io/gorm"
)

func InitializedService(db *gorm.DB) *http.OauthHandler {
	wire.Build(
		repository.NewOauthClientRepository,
		repository.NewOauthAccessTokenRepository,
		repository.NewOauthRefreshTokenRepository,
		http.NewOauthHandler,
		usecase.NewOauthUseCase,
		repository2.NewUserRepository,
		usecase2.NewUserUseCase,
	)
	return &http.OauthHandler{}
}
