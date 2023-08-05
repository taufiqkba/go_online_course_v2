package repository

import (
	"go_online_course_v2/internal/oauth/entity"
	"go_online_course_v2/pkg/response"
	"gorm.io/gorm"
)

type OauthAccessTokenRepository interface {
	Create(entity entity.OauthAccessToken) (*entity.OauthAccessToken, *response.Errors)
	Delete(entity entity.OauthAccessToken) (*entity.OauthAccessToken, *response.Errors)
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

func (repository *oauthAccessTokenRepository) Delete(entity entity.OauthAccessToken) (*entity.OauthAccessToken, *response.Errors) {
	//TODO implement me
	panic("implement me")
}

func (repository *oauthAccessTokenRepository) FindOneByAccessToken(accessToken string) (*entity.OauthAccessToken, *response.Errors) {
	//TODO implement me
	panic("implement me")
}

func NewOauthAccessTokenRepository(db *gorm.DB) OauthAccessTokenRepository {
	return &oauthAccessTokenRepository{db: db}
}
