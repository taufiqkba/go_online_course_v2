package repository

import (
	"go_online_course_v2/internal/forgot_password/entity"
	"go_online_course_v2/pkg/response"
	"gorm.io/gorm"
)

type ForgotPasswordRepository interface {
	Create(entity entity.ForgotPassword) (*entity.ForgotPassword, *response.Errors)
	FindByCode(code string) (*entity.ForgotPassword, *response.Errors)
	Update(entity entity.ForgotPassword) (*entity.ForgotPassword, *response.Errors)
}

type forgotPasswordRepository struct {
	db *gorm.DB
}

func (repository *forgotPasswordRepository) Create(entity entity.ForgotPassword) (*entity.ForgotPassword, *response.Errors) {
	if err := repository.db.Create(&entity).Error; err != nil {
		return nil, &response.Errors{
			Code: 500,
			Err:  err,
		}
	}
	return &entity, nil
}

func (repository *forgotPasswordRepository) FindByCode(code string) (*entity.ForgotPassword, *response.Errors) {
	var forgotPassword entity.ForgotPassword

	if err := repository.db.Where("code = ?", code).First(&forgotPassword).Error; err != nil {
		return nil, &response.Errors{
			Code: 500,
			Err:  err,
		}
	}

	return &forgotPassword, nil
}

func (repository *forgotPasswordRepository) Update(entity entity.ForgotPassword) (*entity.ForgotPassword, *response.Errors) {
	if err := repository.db.Save(&entity).Error; err != nil {
		return nil, &response.Errors{
			Code: 500,
			Err:  nil,
		}
	}
	return &entity, nil
}

func newForgotPasswordRepository(db *gorm.DB) ForgotPasswordRepository {
	return &forgotPasswordRepository{db: db}
}
