package repository

import (
	"go_online_course_v2/internal/admin/entity"
	"go_online_course_v2/pkg/response"
	"go_online_course_v2/pkg/utils"

	"gorm.io/gorm"
)

type AdminRepository interface {
	FindAll(offset int, limit int) []entity.Admin
	FindByID(id int) (*entity.Admin, *response.Errors)
	FindByEmail(email string) (*entity.Admin, *response.Errors)
	Create(entity entity.Admin) (*entity.Admin, *response.Errors)
	Update(entity entity.Admin) (*entity.Admin, *response.Errors)
	Delete(entity entity.Admin) *response.Errors
	TotalCountAdmin() int64
}

type adminRepository struct {
	db *gorm.DB
}

// Create implements AdminRepository.
func (repository *adminRepository) Create(entity entity.Admin) (*entity.Admin, *response.Errors) {
	if err := repository.db.Create(&entity).Error; err != nil {
		return nil, &response.Errors{
			Code: 500,
			Err:  err,
		}
	}
	return &entity, nil
}

// Delete implements AdminRepository.
func (repository *adminRepository) Delete(entity entity.Admin) *response.Errors {
	if err := repository.db.Delete(&entity).Error; err != nil {
		return &response.Errors{
			Code: 500,
			Err:  err,
		}
	}
	return nil
}

// FindAll implements AdminRepository.
func (repository *adminRepository) FindAll(offset int, limit int) []entity.Admin {
	var admins []entity.Admin

	repository.db.Scopes(utils.Paginate(offset, limit)).Find(&admins)
	return admins
}

// FindByEmail implements AdminRepository.
func (repository *adminRepository) FindByEmail(email string) (*entity.Admin, *response.Errors) {
	var admin entity.Admin

	if err := repository.db.Where("email = ?", email).First(&admin).Error; err != nil {
		return nil, &response.Errors{
			Code: 500,
			Err:  err,
		}
	}
	return &admin, nil
}

// FindByID implements AdminRepository.
func (repository *adminRepository) FindByID(id int) (*entity.Admin, *response.Errors) {
	var admin entity.Admin

	if err := repository.db.First(&admin, id).Error; err != nil {
		return nil, &response.Errors{
			Code: 500,
			Err:  err,
		}
	}
	return &admin, nil
}

// TotalCountAdmin implements AdminRepository.
func (repository *adminRepository) TotalCountAdmin() int64 {
	var admin entity.Admin
	var totalAdmin int64

	repository.db.Model(&admin).Count(&totalAdmin)
	return totalAdmin
}

// Update implements AdminRepository.
func (repository *adminRepository) Update(entity entity.Admin) (*entity.Admin, *response.Errors) {
	if err := repository.db.Save(&entity).Error; err != nil {
		return nil, &response.Errors{
			Code: 500,
			Err:  err,
		}
	}
	return &entity, nil
}

func NewAdminRepository(db *gorm.DB) AdminRepository {
	return &adminRepository{db: db}
}
