package main

import (
	"github.com/gin-gonic/gin"
	injector4 "go_online_course_v2/internal/admin/injector"
	injector8 "go_online_course_v2/internal/cart/injector"
	injector7 "go_online_course_v2/internal/discount/injector"
	injector3 "go_online_course_v2/internal/forgot_password/injector"
	injector2 "go_online_course_v2/internal/oauth/injector"
	injector6 "go_online_course_v2/internal/product/injector"
	injector5 "go_online_course_v2/internal/product_category/injector"
	"go_online_course_v2/internal/register/injector"
	"go_online_course_v2/pkg/db/mysql"
)

func main() {
	r := gin.Default()
	db := mysql.DB()

	injector.InitializedService(db).Route(&r.RouterGroup)
	injector2.InitializedService(db).Route(&r.RouterGroup)
	injector3.InitializedService(db).Route(&r.RouterGroup)
	injector4.InitializedService(db).Route(&r.RouterGroup)
	injector5.InitializedService(db).Route(&r.RouterGroup)
	injector6.InitializedService(db).Route(&r.RouterGroup)
	injector7.InitializedService(db).Route(&r.RouterGroup)
	injector8.InitializedService(db).Route(&r.RouterGroup)
	err := r.Run()
	if err != nil {
		return
	}
}
