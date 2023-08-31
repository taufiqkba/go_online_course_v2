//go:build wireinject
// +build wireinject

package injector

import (
	"github.com/google/wire"
	repository7 "go_online_course_v2/internal/admin/repository"
	usecase9 "go_online_course_v2/internal/admin/usecase"
	"go_online_course_v2/internal/cart/repository"
	usecase2 "go_online_course_v2/internal/cart/usecase"
	"go_online_course_v2/internal/dashboard/delivery/http"
	"go_online_course_v2/internal/dashboard/usecase"
	repository2 "go_online_course_v2/internal/discount/repository"
	usecase3 "go_online_course_v2/internal/discount/usecase"
	repository3 "go_online_course_v2/internal/order/repository"
	usecase4 "go_online_course_v2/internal/order/usecase"
	repository4 "go_online_course_v2/internal/order_detail/repository"
	usecase5 "go_online_course_v2/internal/order_detail/usecase"
	usecase6 "go_online_course_v2/internal/payment/usecase"
	repository5 "go_online_course_v2/internal/product/repository"
	usecase7 "go_online_course_v2/internal/product/usecase"
	repository8 "go_online_course_v2/internal/product_category/repository"
	repository6 "go_online_course_v2/internal/user/repository"
	usecase8 "go_online_course_v2/internal/user/usecase"
	"go_online_course_v2/pkg/media/cloudinary"
	"gorm.io/gorm"
)

func InitializedService(db *gorm.DB) *http.DashboardHandler {
	wire.Build(
		http.NewDashboardHandler,
		usecase.NewDashboardUseCase,
		repository.NewCartRepository,
		usecase2.NewCartUseCase,
		repository2.NewDiscountRepository,
		usecase3.NewDiscountUseCase,
		repository3.NewOrderRepository,
		usecase4.NewOrderUseCase,
		repository4.NewOrderDetailRepository,
		usecase5.NewOrderDetailUseCase,
		usecase6.NewPaymentUseCase,
		repository5.NewProductRepository,
		usecase7.NewProductUseCase,
		repository6.NewUserRepository,
		usecase8.NewUserUseCase,
		repository7.NewAdminRepository,
		usecase9.NewAdminUseCase,
		cloudinary.NewMediaUseCase,
		repository8.NewProductCategoryRepository,
	)

	return &http.DashboardHandler{}
}
