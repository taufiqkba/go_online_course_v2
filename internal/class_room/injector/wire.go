//go:build wireinject
// +build wireinject

package injector

import (
	"github.com/google/wire"
	"go_online_course_v2/internal/class_room/delivery/http"
	"go_online_course_v2/internal/class_room/repository"
	"go_online_course_v2/internal/class_room/usecase"
	"gorm.io/gorm"
)

func InitializedService(db *gorm.DB) *http.ClassRoomHandler {
	wire.Build(
		repository.NewClassRoomRepository,
		usecase.NewClassRoomUseCase,
		http.NewClassRoomHandler,
	)
	return &http.ClassRoomHandler{}
}
