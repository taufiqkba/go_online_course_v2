package dto

import (
	entity3 "go_online_course_v2/internal/class_room/entity"
	entity2 "go_online_course_v2/internal/product/entity"
	"go_online_course_v2/internal/user/entity"
	"gorm.io/gorm"
	"time"
)

type ClassRoomResponseBody struct {
	ID        int64            `json:"id"`
	User      *entity.User     `json:"user"`
	Product   *entity2.Product `json:"product"`
	CreatedBy *entity.User     `json:"created_by"`
	UpdatedBy *entity.User     `json:"updated_by"`
	CreatedAt *time.Time       `json:"created_at"`
	UpdatedAt *time.Time       `json:"updated_at"`
	DeletedAt gorm.DeletedAt   `json:"deleted_at"`
}

func CreateClassRoomResponse(entity entity3.ClassRoom) ClassRoomResponseBody {
	return ClassRoomResponseBody{
		ID:        entity.ID,
		User:      entity.User,
		Product:   entity.Product,
		CreatedBy: entity.CreatedBy,
		UpdatedBy: entity.UpdatedBy,
		CreatedAt: entity.CreatedAt,
		UpdatedAt: entity.UpdatedAt,
		DeletedAt: entity.DeletedAt,
	}
}

type ClassRoomListResponse []ClassRoomResponseBody

func CreateClassRoomListResponse(entity []entity3.ClassRoom) ClassRoomListResponse {
	classRoomResp := ClassRoomListResponse{}

	for _, classRoom := range entity {
		classRoom.Product.VideoLink = classRoom.Product.Video

		classRoomResponse := CreateClassRoomResponse(classRoom)
		classRoomResp = append(classRoomResp, classRoomResponse)
	}
	return classRoomResp
}
