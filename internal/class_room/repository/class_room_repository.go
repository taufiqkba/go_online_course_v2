package repository

import (
	"go_online_course_v2/internal/class_room/entity"
	"go_online_course_v2/pkg/response"
	"go_online_course_v2/pkg/utils"
	"gorm.io/gorm"
)

type ClassRoomRepository interface {
	FindAllByUserID(userID int, offset int, limit int) []entity.ClassRoom
	FindOneByUserIDAndProductID(userID int, productID int) (*entity.ClassRoom, *response.Errors)
	Create(entity entity.ClassRoom) (*entity.ClassRoom, *response.Errors)
}

type classRoomRepository struct {
	db *gorm.DB
}

func (repository *classRoomRepository) FindAllByUserID(userID int, offset int, limit int) []entity.ClassRoom {
	var classRooms []entity.ClassRoom

	repository.db.Scopes(utils.Paginate(offset, limit)).
		Preload("Product.ProductCategory").
		Where("user_id = ?", userID).
		Find(&classRooms)
	return classRooms
}

func (repository *classRoomRepository) FindOneByUserIDAndProductID(userID int, productID int) (*entity.ClassRoom, *response.Errors) {
	var classRoom entity.ClassRoom

	if err := repository.db.
		Preload("Product.ProductCategory").
		Where("user_id = ?", userID).
		Where("product_id = ?", productID).
		First(&classRoom).Error; err != nil {
		return nil, &response.Errors{
			Code: 500,
			Err:  err,
		}
	}
	return &classRoom, nil
}

func (repository *classRoomRepository) Create(entity entity.ClassRoom) (*entity.ClassRoom, *response.Errors) {
	if err := repository.db.Create(&entity).Error; err != nil {
		return nil, &response.Errors{
			Code: 500,
			Err:  err,
		}
	}
	return &entity, nil
}

func NewClassRoomRepository(db *gorm.DB) ClassRoomRepository {
	return &classRoomRepository{db: db}
}
