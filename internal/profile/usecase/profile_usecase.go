package usecase

import (
	usecase2 "go_online_course_v2/internal/oauth/usecase"
	"go_online_course_v2/internal/profile/dto"
	dto2 "go_online_course_v2/internal/user/dto"
	"go_online_course_v2/internal/user/entity"
	"go_online_course_v2/internal/user/usecase"
	"go_online_course_v2/pkg/response"
)

type ProfileUseCase interface {
	FindProfile(id int) (*dto.ProfileResponseBody, *response.Errors)
	Update(id int, dto dto2.UserUpdateRequestBody) (*entity.User, *response.Errors)
	Deactivated(id int) *response.Errors
	Logout(accessToken string) *response.Errors
}

type profileUseCase struct {
	userUseCase  usecase.UserUseCase
	oauthUseCase usecase2.OauthUseCase
}

func (useCase *profileUseCase) FindProfile(id int) (*dto.ProfileResponseBody, *response.Errors) {
	user, err := useCase.userUseCase.FindOneByID(id)

	if err != nil {
		return nil, err
	}

	userResponse := dto.CreateProfileResponse(*user)
	return &userResponse, nil
}

func (useCase *profileUseCase) Update(id int, dto dto2.UserUpdateRequestBody) (*entity.User, *response.Errors) {
	//	find data by id
	user, err := useCase.userUseCase.FindOneByID(id)
	if err != nil {
		return nil, err
	}

	updateUser, err := useCase.userUseCase.Update(int(user.ID), dto)
	if err != nil {
		return nil, err
	}
	return updateUser, nil
}

func (useCase *profileUseCase) Deactivated(id int) *response.Errors {
	user, err := useCase.userUseCase.FindOneByID(id)
	if err != nil {
		return err
	}

	err = useCase.userUseCase.Delete(int(user.ID))
	if err != nil {
		return err
	}
	return nil
}

func (useCase *profileUseCase) Logout(accessToken string) *response.Errors {
	return useCase.oauthUseCase.Logout(accessToken)
}

func NewProfileUseCase(userUseCase usecase.UserUseCase, oauthUseCase usecase2.OauthUseCase) ProfileUseCase {
	return &profileUseCase{userUseCase: userUseCase, oauthUseCase: oauthUseCase}
}
