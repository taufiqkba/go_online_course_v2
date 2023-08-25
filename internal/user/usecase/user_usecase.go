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
	Update(id int, dto dto.UserUpdateRequestBody) (*entity.User, *response.Errors)
	Delete(id int) *response.Errors
	TotalCountUser() int64
}

type userUseCase struct {
	repository repository.UserRepository
}

func (useCase *userUseCase) FindAll(offset int, limit int) []entity.User {
	return useCase.repository.FindAll(offset, limit)
}

// FindByEmail implements UserUseCase
func (useCase *userUseCase) FindByEmail(email string) (*entity.User, *response.Errors) {
	return useCase.repository.FindByEmail(email)
}

func (useCase *userUseCase) FindOneByID(id int) (*entity.User, *response.Errors) {
	return useCase.repository.FindOneByID(id)
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
		CodeVerified: utils.RandNumber(12),
	}

	if dto.CreatedBy != nil {
		user.CreatedByID = dto.CreatedBy
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
	return useCase.repository.FindOneByCodeVerified(codeVerified)
}

func (useCase *userUseCase) Update(id int, dto dto.UserUpdateRequestBody) (*entity.User, *response.Errors) {
	//	find user by id
	user, err := useCase.repository.FindOneByID(id)

	if err != nil {
		return nil, err
	}

	if dto.Email != nil {
		if user.Email != *dto.Email {
			user.Email = *dto.Email
		}
	}

	if dto.Name != nil {
		user.Name = *dto.Name
	}

	//	check
	if dto.Password != nil {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(*dto.Password), bcrypt.DefaultCost)
		if err != nil {
			return nil, &response.Errors{
				Code: 500,
				Err:  err,
			}
		}
		user.Password = string(hashedPassword)
	}

	if dto.UpdatedBy != nil {
		user.UpdatedByID = dto.UpdatedBy
	}

	updateUser, err := useCase.repository.Update(*user)
	if err != nil {
		return nil, err
	}

	return updateUser, nil
}

func (useCase *userUseCase) Delete(id int) *response.Errors {
	user, err := useCase.repository.FindOneByID(id)
	if err != nil {
		return err
	}
	err = useCase.repository.Delete(*user)
	if err != nil {
		return err
	}
	return nil
}

func (useCase *userUseCase) TotalCountUser() int64 {
	//TODO implement me
	panic("implement me")
}

func NewUserUseCase(repository repository.UserRepository) UserUseCase {
	return &userUseCase{repository: repository}
}
