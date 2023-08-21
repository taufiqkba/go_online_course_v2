//go:build wireinject
// +build wireinject

package injector

import (
	"github.com/google/wire"
	repository4 "go_online_course_v2/internal/cart/repository"
	usecase5 "go_online_course_v2/internal/cart/usecase"
	"go_online_course_v2/internal/class_room/repository"
	usecase2 "go_online_course_v2/internal/class_room/usecase"
	repository5 "go_online_course_v2/internal/discount/repository"
	usecase6 "go_online_course_v2/internal/discount/usecase"
	repository2 "go_online_course_v2/internal/order/repository"
	usecase3 "go_online_course_v2/internal/order/usecase"
	repository3 "go_online_course_v2/internal/order_detail/repository"
	usecase4 "go_online_course_v2/internal/order_detail/usecase"
	usecase7 "go_online_course_v2/internal/payment/usecase"
	repository6 "go_online_course_v2/internal/product/repository"
	usecase8 "go_online_course_v2/internal/product/usecase"
	repository7 "go_online_course_v2/internal/product_category/repository"
	"go_online_course_v2/internal/webhook/delivery/http"
	"go_online_course_v2/internal/webhook/usecase"
	"go_online_course_v2/pkg/media/cloudinary"
	"gorm.io/gorm"
)

func InitializedService(db *gorm.DB) *http.WebHookHandler {
	wire.Build(
		http.NewWebHookHandler,
		usecase.NewWebHookUseCase,
		repository.NewClassRoomRepository,
		usecase2.NewClassRoomUseCase,
		repository2.NewOrderRepository,
		usecase3.NewOrderUseCase,
		repository3.NewOrderDetailRepository,
		usecase4.NewOrderDetailUseCase,
		repository4.NewCartRepository,
		usecase5.NewCartUseCase,
		repository5.NewDiscountRepository,
		usecase6.NewDiscountUseCase,
		usecase7.NewPaymentUseCase,
		repository6.NewProductRepository,
		usecase8.NewProductUseCase,
		repository7.NewProductCategoryRepository,
		cloudinary.NewMediaUseCase,
	)

	return &http.WebHookHandler{}
}
