// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package injector

import (
	repository2 "go_online_course_v2/internal/cart/repository"
	"go_online_course_v2/internal/cart/usecase"
	repository3 "go_online_course_v2/internal/discount/repository"
	usecase2 "go_online_course_v2/internal/discount/usecase"
	"go_online_course_v2/internal/order/delivery/http"
	"go_online_course_v2/internal/order/repository"
	usecase6 "go_online_course_v2/internal/order/usecase"
	repository4 "go_online_course_v2/internal/order_detail/repository"
	usecase3 "go_online_course_v2/internal/order_detail/usecase"
	usecase4 "go_online_course_v2/internal/payment/usecase"
	repository5 "go_online_course_v2/internal/product/repository"
	usecase5 "go_online_course_v2/internal/product/usecase"
	repository6 "go_online_course_v2/internal/product_category/repository"
	"go_online_course_v2/pkg/media/cloudinary"
	"gorm.io/gorm"
)

// Injectors from wire.go:

func InitializedService(db *gorm.DB) *http.OrderHandler {
	orderRepository := repository.NewOrderRepository(db)
	cartRepository := repository2.NewCartRepository(db)
	cartUseCase := usecase.NewCartUseCase(cartRepository)
	discountRepository := repository3.NewDiscountRepository(db)
	discountUseCase := usecase2.NewDiscountUseCase(discountRepository)
	orderDetailRepository := repository4.NewOrderDetailRepository(db)
	orderDetailUseCase := usecase3.NewOrderDetailUseCase(orderDetailRepository)
	paymentUseCase := usecase4.NewPaymentUseCase()
	productRepository := repository5.NewProductRepository(db)
	productCategoryRepository := repository6.NewProductCategoryRepository(db)
	media := cloudinary.NewMediaUseCase()
	productUseCase := usecase5.NewProductUseCase(productRepository, productCategoryRepository, media)
	orderUseCase := usecase6.NewOrderUseCase(orderRepository, cartUseCase, discountUseCase, orderDetailUseCase, paymentUseCase, productUseCase)
	orderHandler := http.NewOrderHandler(orderUseCase)
	return orderHandler
}
