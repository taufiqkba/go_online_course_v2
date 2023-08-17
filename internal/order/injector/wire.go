//go:build wireinject
// +build wireinject

package injector

import (
	"github.com/google/wire"
	"go_online_course_v2/internal/cart/repository"
	"go_online_course_v2/internal/cart/usecase"
	repository2 "go_online_course_v2/internal/discount/repository"
	usecase2 "go_online_course_v2/internal/discount/usecase"
	"go_online_course_v2/internal/order/delivery/http"
	repository5 "go_online_course_v2/internal/order/repository"
	usecase6 "go_online_course_v2/internal/order/usecase"
	repository3 "go_online_course_v2/internal/order_detail/repository"
	usecase3 "go_online_course_v2/internal/order_detail/usecase"
	usecase5 "go_online_course_v2/internal/payment/usecase"
	repository4 "go_online_course_v2/internal/product/repository"
	usecase4 "go_online_course_v2/internal/product/usecase"
	repository6 "go_online_course_v2/internal/product_category/repository"
	"go_online_course_v2/pkg/media/cloudinary"
	"gorm.io/gorm"
)

func InitializedService(db *gorm.DB) *http.OrderHandler {
	wire.Build(
		repository.NewCartRepository,
		usecase.NewCartUseCase,
		repository2.NewDiscountRepository,
		usecase2.NewDiscountUseCase,
		repository3.NewOrderDetailRepository,
		usecase3.NewOrderDetailUseCase,
		repository4.NewProductRepository,
		repository6.NewProductCategoryRepository,
		usecase4.NewProductUseCase,
		usecase5.NewPaymentUseCase,
		cloudinary.NewMediaUseCase,
		repository5.NewOrderRepository,
		usecase6.NewOrderUseCase,
		http.NewOrderHandler,
	)

	return &http.OrderHandler{}
}
