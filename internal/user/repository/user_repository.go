package repository

import (
	"go_online_course_v2/internal/user/entity"
	"go_online_course_v2/pkg/response"
	"gorm.io/gorm"
)

type UserRepository interface {
	FindAll(offset int, limit int) []entity.User
	FindOneByID(id int) (*entity.User, *response.Errors)
	FindByEmail(email string) (*entity.User, *response.Errors)
	Create(entity entity.User) (*entity.User, *response.Errors)
	FindOneByCodeVerified(codeVerified string) (*entity.User, *response.Errors)
	Update(id int, entity entity.User) (*entity.User, *response.Errors)
	Delete(id int, entity entity.User) (*entity.User, *response.Errors)
	TotalCountUser() int64
}

type userRepository struct {
	db *gorm.DB
}

func (repository *userRepository) FindAll(offset int, limit int) []entity.User {
	//TODO implement me
	panic("implement me")
}

func (repository *userRepository) FindOneByID(id int) (*entity.User, *response.Errors) {
	//TODO implement me
	panic("implement me")
}

func (repository *userRepository) FindByEmail(email string) (*entity.User, *response.Errors) {
	//TODO implement me
	panic("implement me")
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
	//TODO implement me
	panic("implement me")
}

func (repository *userRepository) Update(id int, entity entity.User) (*entity.User, *response.Errors) {
	//TODO implement me
	panic("implement me")
}

func (repository *userRepository) Delete(id int, entity entity.User) (*entity.User, *response.Errors) {
	//TODO implement me
	panic("implement me")
}

func (repository *userRepository) TotalCountUser() int64 {
	//TODO implement me
	panic("implement me")
}

func newUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}
