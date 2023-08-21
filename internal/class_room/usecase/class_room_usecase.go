package usecase

import (
	"errors"
	"go_online_course_v2/internal/class_room/dto"
	"go_online_course_v2/internal/class_room/entity"
	"go_online_course_v2/internal/class_room/repository"
	"go_online_course_v2/pkg/response"
	"gorm.io/gorm"
)

type ClassRoomUseCase interface {
	FindAllByUserID(userID int, offset int, limit int) dto.ClassRoomListResponse
	FindOneByUserIDAndProductID(userID int, productID int) (*dto.ClassRoomResponseBody, *response.Errors)
	Create(dto dto.ClassRoomRequestBody) (*entity.ClassRoom, *response.Errors)
}

type classRoomUseCase struct {
	repository repository.ClassRoomRepository
}

func (useCase *classRoomUseCase) FindAllByUserID(userID int, offset int, limit int) dto.ClassRoomListResponse {
	classRooms := useCase.repository.FindAllByUserID(userID, offset, limit)

	classRoomResp := dto.CreateClassRoomListResponse(classRooms)
	return classRoomResp
}

func (useCase *classRoomUseCase) FindOneByUserIDAndProductID(userID int, productID int) (*dto.ClassRoomResponseBody, *response.Errors) {
	classRoom, _ := useCase.repository.FindOneByUserIDAndProductID(userID, productID)

	classRoomResp := dto.CreateClassRoomResponse(*classRoom)
	return &classRoomResp, nil
}

func (useCase *classRoomUseCase) Create(dto dto.ClassRoomRequestBody) (*entity.ClassRoom, *response.Errors) {
	//	validate product and user
	dataClassRoom, err := useCase.repository.FindOneByUserIDAndProductID(int(dto.UserID), int(dto.ProductID))
	if err != nil && !errors.Is(err.Err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	if dataClassRoom != nil {
		return nil, &response.Errors{
			Code: 400,
			Err:  errors.New("you've sign to this class room"),
		}
	}

	classRoom := entity.ClassRoom{
		UserID:      dto.UserID,
		ProductID:   &dto.ProductID,
		CreatedByID: &dto.UserID,
	}

	data, err := useCase.repository.Create(classRoom)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func NewClassRoomUseCase(repository repository.ClassRoomRepository) ClassRoomUseCase {
	return &classRoomUseCase{repository: repository}
}
