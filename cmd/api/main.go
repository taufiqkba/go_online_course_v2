package main

import (
	"github.com/gin-gonic/gin"
	injector2 "go_online_course_v2/internal/oauth/injector"
	"go_online_course_v2/internal/register/injector"
	"go_online_course_v2/pkg/db/mysql"
)

func main() {
	r := gin.Default()
	db := mysql.DB()

	injector.InitializedService(db).Route(&r.RouterGroup)
	injector2.InitializedService(db).Route(&r.RouterGroup)
	err := r.Run()
	if err != nil {
		return
	}
}
