//go:build wireinject
// +build wireinject

package injector

import (
	"github.com/google/wire"
	"go_online_course_v2/internal/product/delivery/http"
	"go_online_course_v2/internal/product/repository"
	"go_online_course_v2/internal/product/usecase"
	repository2 "go_online_course_v2/internal/product_category/repository"
	"go_online_course_v2/pkg/media/cloudinary"
	"gorm.io/gorm"
)

func InitializedService(db *gorm.DB) *http.ProductHandler {
	wire.Build(
		repository2.NewProductCategoryRepository,
		repository.NewProductRepository,
		usecase.NewProductUseCase,
		http.NewProductHandler,
		cloudinary.NewMediaUseCase,
	)
	return &http.ProductHandler{}
}
