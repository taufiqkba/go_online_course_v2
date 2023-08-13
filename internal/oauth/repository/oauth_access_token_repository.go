package repository

import (
	"go_online_course_v2/internal/oauth/entity"
	"go_online_course_v2/pkg/response"
	"gorm.io/gorm"
)

type OauthAccessTokenRepository interface {
	Create(entity entity.OauthAccessToken) (*entity.OauthAccessToken, *response.Errors)
	Delete(entity entity.OauthAccessToken) *response.Errors
	FindOneByAccessToken(accessToken string) (*entity.OauthAccessToken, *response.Errors)
}

type oauthAccessTokenRepository struct {
	db *gorm.DB
}

func (repository *oauthAccessTokenRepository) Create(entity entity.OauthAccessToken) (*entity.OauthAccessToken, *response.Errors) {
	if err := repository.db.Create(&entity).Error; err != nil {
		return nil, &response.Errors{
			Code: 500,
			Err:  err,
		}
	}
	return &entity, nil
}

func (repository *oauthAccessTokenRepository) Delete(entity entity.OauthAccessToken) *response.Errors {
	if err := repository.db.Delete(&entity).Error; err != nil {
		return &response.Errors{
			Code: 500,
			Err:  err,
		}
	}
	return nil
}

func (repository *oauthAccessTokenRepository) FindOneByAccessToken(accessToken string) (*entity.OauthAccessToken, *response.Errors) {
	var oauthAccessToken entity.OauthAccessToken

	if err := repository.db.
		Where("token = ?", accessToken).
		First(&oauthAccessToken).
		Error; err != nil {
		return nil, &response.Errors{
			Code: 500,
			Err:  err,
		}
	}
	return &oauthAccessToken, nil
}

func NewOauthAccessTokenRepository(db *gorm.DB) OauthAccessTokenRepository {
	return &oauthAccessTokenRepository{db: db}
}
