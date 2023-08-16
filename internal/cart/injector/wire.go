//go:build wireinject
// +build wireinject

package injector

import (
	"github.com/google/wire"
	"go_online_course_v2/internal/cart/delivery/http"
	"go_online_course_v2/internal/cart/repository"
	"go_online_course_v2/internal/cart/usecase"
	"gorm.io/gorm"
)

func InitializedService(db *gorm.DB) *http.CartHandler {
	wire.Build(
		repository.NewCartRepository,
		usecase.NewCartUseCase,
		http.NewCartHandler,
	)
	return &http.CartHandler{}
}
