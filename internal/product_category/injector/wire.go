//go:build wireinject
// +build wireinject

package injector

import (
	"github.com/google/wire"
	"go_online_course_v2/internal/product_category/delivery/http"
	"go_online_course_v2/internal/product_category/repository"
	"go_online_course_v2/internal/product_category/usecase"
	"go_online_course_v2/pkg/media/cloudinary"
	"gorm.io/gorm"
)

func InitializedService(db *gorm.DB) *http.ProductCategoryHandler {
	wire.Build(
		repository.NewProductCategoryRepository,
		usecase.NewProductCategoryUseCase,
		http.NewProductCategoryHandler,
		cloudinary.NewMediaUseCase,
	)
	return &http.ProductCategoryHandler{}
}
