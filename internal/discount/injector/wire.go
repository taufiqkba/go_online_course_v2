//go:build wireinject
// +build wireinject

package injector

import (
	"github.com/google/wire"
	"go_online_course_v2/internal/discount/delivery"
	"go_online_course_v2/internal/discount/repository"
	"go_online_course_v2/internal/discount/usecase"
	"gorm.io/gorm"
)

func InitializedService(db *gorm.DB) *delivery.DiscountHandler {
	wire.Build(
		repository.NewDiscountRepository,
		usecase.NewDiscountUseCase,
		delivery.NewDiscountHandler,
	)
	return &delivery.DiscountHandler{}
}
