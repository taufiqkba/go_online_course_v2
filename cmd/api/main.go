package main

import (
	"github.com/gin-gonic/gin"
	"go_online_course_v2/internal/oauth/delivery"
	repository2 "go_online_course_v2/internal/oauth/repository"
	usecase3 "go_online_course_v2/internal/oauth/usecase"
	"go_online_course_v2/internal/register/delivery/http"
	usecase2 "go_online_course_v2/internal/register/usecase"
	"go_online_course_v2/internal/user/repository"
	"go_online_course_v2/internal/user/usecase"
	"go_online_course_v2/pkg/db/mysql"
)

func main() {
	r := gin.Default()
	db := mysql.DB()

	userRepository := repository.NewUserRepository(db)
	userUseCase := usecase.NewUserUseCase(userRepository)

	registerUseCase := usecase2.NewRegisterUseCase(userUseCase)
	http.NewRegisterHandler(registerUseCase).Route(&r.RouterGroup)

	oauthClientRepository := repository2.NewOauthClientRepository(db)
	oauthAccessTokenRepository := repository2.NewOauthAccessTokenRepository(db)
	oauthRefreshTokenRepository := repository2.NewOauthRefreshTokenRepository(db)
	oauthUseCase := usecase3.NewOauthUseCase(
		oauthClientRepository,
		oauthAccessTokenRepository,
		oauthRefreshTokenRepository,
		userUseCase,
	)

	delivery.NewOauthHandler(oauthUseCase).Route(&r.RouterGroup)

	err := r.Run()
	if err != nil {
		return
	}
}
