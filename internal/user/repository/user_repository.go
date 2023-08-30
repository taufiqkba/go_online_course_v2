package repository

import (
	"go_online_course_v2/internal/user/entity"
	"go_online_course_v2/pkg/response"
	"go_online_course_v2/pkg/utils"
	"gorm.io/gorm"
)

type UserRepository interface {
	FindAll(offset int, limit int) []entity.User
	FindOneByID(id int) (*entity.User, *response.Errors)
	FindByEmail(email string) (*entity.User, *response.Errors)
	Create(entity entity.User) (*entity.User, *response.Errors)
	FindOneByCodeVerified(codeVerified string) (*entity.User, *response.Errors)
	Update(entity entity.User) (*entity.User, *response.Errors)
	Delete(entity entity.User) *response.Errors
	TotalCountUser() int64
}

type userRepository struct {
	db *gorm.DB
}

func (repository *userRepository) FindAll(offset int, limit int) []entity.User {
	var users []entity.User

	repository.db.Scopes(utils.Paginate(offset, limit)).Find(&users)
	return users
}

func (repository *userRepository) FindOneByID(id int) (*entity.User, *response.Errors) {
	var user entity.User
	if err := repository.db.First(&user, id).Error; err != nil {
		return nil, &response.Errors{
			Code: 500,
			Err:  nil,
		}
	}
	return &user, nil
}

func (repository *userRepository) FindByEmail(email string) (*entity.User, *response.Errors) {
	var user entity.User

	if err := repository.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, &response.Errors{
			Code: 500,
			Err:  err,
		}
	}
	return &user, nil
}

func (repository *userRepository) Create(entity entity.User) (*entity.User, *response.Errors) {
	if err := repository.db.Create(&entity).Error; err != nil {
		return nil, &response.Errors{
			Code: 500,
			Err:  err,
		}
	}
	return &entity, nil
}

func (repository *userRepository) FindOneByCodeVerified(codeVerified string) (*entity.User, *response.Errors) {
	var user entity.User
	if err := repository.db.
		Where("code_verified = ?", codeVerified).
		First(&user).
		Error; err != nil {
		return nil, &response.Errors{
			Code: 500,
			Err:  err,
		}
	}
	return &user, nil
}

func (repository *userRepository) Update(entity entity.User) (*entity.User, *response.Errors) {
	if err := repository.db.Save(&entity).Error; err != nil {
		return nil, &response.Errors{
			Code: 500,
			Err:  err,
		}
	}
	return &entity, nil
}

func (repository *userRepository) Delete(entity entity.User) *response.Errors {
	if err := repository.db.Delete(&entity).Error; err != nil {
		return &response.Errors{
			Code: 500,
			Err:  err,
		}
	}
	return nil
}

func (repository *userRepository) TotalCountUser() int64 {
	var user entity.User
	var totalUser int64

	repository.db.Model(&user).Count(&totalUser)
	return totalUser
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}
