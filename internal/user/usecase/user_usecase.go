package usecase

import (
	"errors"
	"go_online_course_v2/internal/user/dto"
	"go_online_course_v2/internal/user/entity"
	"go_online_course_v2/internal/user/repository"
	"go_online_course_v2/pkg/response"
	"go_online_course_v2/pkg/utils"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserUseCase interface {
	FindAll(offset int, limit int) []entity.User
	FindByEmail(email string) (*entity.User, *response.Errors)
	FindOneByID(id int) (*entity.User, *response.Errors)
	Create(dto dto.UserRequestBody) (*entity.User, *response.Errors)
	FindOneByCodeVerified(codeVerified string) (*entity.User, *response.Errors)
	Update(id int, dto dto.UserRequestBody) (*entity.User, *response.Errors)
	Delete(id int) *response.Errors
	TotalCountUser() int64
}

type userUseCase struct {
	repository repository.UserRepository
}

func (useCase *userUseCase) FindAll(offset int, limit int) []entity.User {
	//TODO implement me
	panic("implement me")
}

// FindByEmail implements UserUseCase
func (useCase *userUseCase) FindByEmail(email string) (*entity.User, *response.Errors) {
	//	TODO IMPLEMENT Me
	panic("implement me")
}

func (useCase *userUseCase) FindOneByID(id int) (*entity.User, *response.Errors) {
	//TODO implement me
	panic("implement me")
}

// Create implements UserUseCase
func (useCase *userUseCase) Create(dto dto.UserRequestBody) (*entity.User, *response.Errors) {
	//find by email
	checkUser, err := useCase.repository.FindByEmail(dto.Email)
	if err != nil && !errors.Is(err.Err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	if checkUser != nil {
		return nil, &response.Errors{
			Code: 409,
			Err:  errors.New("email has been registered"),
		}
	}

	//	hash password
	hashedPassword, errHashedPassword := bcrypt.GenerateFromPassword([]byte(dto.Password), bcrypt.DefaultCost)
	if errHashedPassword != nil {
		return nil, &response.Errors{
			Code: 500,
			Err:  errHashedPassword,
		}
	}

	user := entity.User{
		Name:         dto.Name,
		Email:        dto.Email,
		Password:     string(hashedPassword),
		CodeVerified: utils.RandString(32),
	}

	dataUser, err := useCase.repository.Create(user)
	if err != nil {
		return nil, &response.Errors{
			Code: 500,
			Err:  errHashedPassword,
		}
	}
	return dataUser, nil
}

func (useCase *userUseCase) FindOneByCodeVerified(codeVerified string) (*entity.User, *response.Errors) {
	//TODO implement me
	panic("implement me")
}

func (useCase *userUseCase) Update(id int, dto dto.UserRequestBody) (*entity.User, *response.Errors) {
	//TODO implement me
	panic("implement me")
}

func (useCase *userUseCase) Delete(id int) *response.Errors {
	//TODO implement me
	panic("implement me")
}

func (useCase *userUseCase) TotalCountUser() int64 {
	//TODO implement me
	panic("implement me")
}

func NewUserUseCase(repository repository.UserRepository) UserUseCase {
	return &userUseCase{repository: repository}
}
