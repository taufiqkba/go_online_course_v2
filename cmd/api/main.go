package main

import (
	"github.com/gin-gonic/gin"
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
	err := r.Run()
	if err != nil {
		return
	}
}
