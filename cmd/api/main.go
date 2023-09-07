package main

import (
	injector4 "go_online_course_v2/internal/admin/injector"
	injector8 "go_online_course_v2/internal/cart/injector"
	injector10 "go_online_course_v2/internal/class_room/injector"
	injector13 "go_online_course_v2/internal/dashboard/injector"
	injector7 "go_online_course_v2/internal/discount/injector"
	injector3 "go_online_course_v2/internal/forgot_password/injector"
	injector2 "go_online_course_v2/internal/oauth/injector"
	injector9 "go_online_course_v2/internal/order/injector"
	injector6 "go_online_course_v2/internal/product/injector"
	injector5 "go_online_course_v2/internal/product_category/injector"
	injector14 "go_online_course_v2/internal/profile/injector"
	"go_online_course_v2/internal/register/injector"
	injector12 "go_online_course_v2/internal/user/injector"
	injector15 "go_online_course_v2/internal/verification_email/injector"
	injector11 "go_online_course_v2/internal/webhook/injector"
	"go_online_course_v2/pkg/db/mysql"

	"github.com/gin-gonic/gin"
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
	injector9.InitializedService(db).Route(&r.RouterGroup)
	injector10.InitializedService(db).Route(&r.RouterGroup)
	injector11.InitializedService(db).Route(&r.RouterGroup)
	injector12.InitializedService(db).Route(&r.RouterGroup)
	injector13.InitializedService(db).Route(&r.RouterGroup)
	injector14.InitializedService(db).Route(&r.RouterGroup)
	injector15.InitializedService(db).Route(&r.RouterGroup)
	err := r.Run()
	if err != nil {
		return
	}
}
