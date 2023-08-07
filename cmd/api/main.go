package main

import (
	"github.com/gin-gonic/gin"
	"go_online_course_v2/internal/oauth/delivery"
	repository2 "go_online_course_v2/internal/oauth/repository"
	usecase3 "go_online_course_v2/internal/oauth/usecase"
	"go_online_course_v2/internal/register/injector"
	"go_online_course_v2/pkg/db/mysql"
)

func main() {
	r := gin.Default()
	db := mysql.DB()

	injector.InitializedService(db).Route(&r.RouterGroup)

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
