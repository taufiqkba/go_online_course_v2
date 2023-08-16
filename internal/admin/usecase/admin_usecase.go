package usecase

import (
	"go_online_course_v2/internal/admin/dto"
	"go_online_course_v2/internal/admin/entity"
	"go_online_course_v2/internal/admin/repository"
	"go_online_course_v2/pkg/response"
	"golang.org/x/crypto/bcrypt"
)

type AdminUseCase interface {
	FindAll(offset int, limit int) []entity.Admin
	FindByID(id int) (*entity.Admin, *response.Errors)
	FindByEmail(email string) (*entity.Admin, *response.Errors)
	Create(dto dto.AdminRequestBody) (*entity.Admin, *response.Errors)
	Update(id int, dto dto.AdminRequestBody) (*entity.Admin, *response.Errors)
	Delete(id int) *response.Errors
	TotalCountAdmin() int64
}

type adminUseCase struct {
	repository repository.AdminRepository
}

// Create implements AdminUseCase.
func (useCase *adminUseCase) Create(dto dto.AdminRequestBody) (*entity.Admin, *response.Errors) {
	//	hashed password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(*dto.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, &response.Errors{
			Code: 500,
			Err:  err,
		}
	}

	dataAdmin := entity.Admin{
		Name:     dto.Name,
		Email:    dto.Email,
		Password: string(hashedPassword),
	}

	if dto.CreatedBy != nil {
		dataAdmin.CreatedByID = dto.CreatedBy
	}

	admin, errCreateAdmin := useCase.repository.Create(dataAdmin)
	if errCreateAdmin != nil {
		return nil, errCreateAdmin
	}
	return admin, nil
}

// Delete implements AdminUseCase.
func (useCase *adminUseCase) Delete(id int) *response.Errors {
	//find by id
	admin, err := useCase.repository.FindByID(id)

	if err != nil {
		return err
	}

	if err := useCase.repository.Delete(*admin); err != nil {
		return err
	}
	return nil
}

// FindAll implements AdminUseCase.
func (useCase *adminUseCase) FindAll(offset int, limit int) []entity.Admin {
	return useCase.repository.FindAll(offset, limit)
}

// FindByEmail implements AdminUseCase.
func (useCase *adminUseCase) FindByEmail(email string) (*entity.Admin, *response.Errors) {
	return useCase.repository.FindByEmail(email)
}

// FindByID implements AdminUseCase.
func (useCase *adminUseCase) FindByID(id int) (*entity.Admin, *response.Errors) {
	return useCase.repository.FindByID(id)
}

// TotalCountAdmin implements AdminUseCase.
func (useCase *adminUseCase) TotalCountAdmin() int64 {
	panic("unimplemented")
}

// Update implements AdminUseCase.
func (useCase *adminUseCase) Update(id int, dto dto.AdminRequestBody) (*entity.Admin, *response.Errors) {
	//	find admin by id
	admin, err := useCase.repository.FindByID(id)
	if err != nil {
		return nil, err
	}
	admin.Name = dto.Name
	admin.Email = dto.Email

	if dto.Password != nil {
		hashedPassword, erHashedPassword := bcrypt.GenerateFromPassword([]byte(*dto.Password), bcrypt.DefaultCost)
		if erHashedPassword != nil {
			return nil, &response.Errors{
				Code: 500,
				Err:  erHashedPassword,
			}
		}
		admin.Password = string(hashedPassword)
	}

	if dto.UpdatedBy != nil {
		admin.UpdatedByID = dto.UpdatedBy
	}

	updateAdmin, err := useCase.repository.Update(*admin)
	if err != nil {
		return nil, err
	}

	return updateAdmin, nil
}

func NewAdminUseCase(repository repository.AdminRepository) AdminUseCase {
	return &adminUseCase{repository: repository}
}
