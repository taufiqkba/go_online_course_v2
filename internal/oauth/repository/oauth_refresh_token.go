package repository

import (
	"go_online_course_v2/internal/oauth/entity"
	"go_online_course_v2/pkg/response"
	"gorm.io/gorm"
)

type OauthRefreshTokenRepository interface {
	Create(entity entity.OauthRefreshToken) (*entity.OauthRefreshToken, *response.Errors)
	Delete(entity entity.OauthRefreshToken) *response.Errors
	FindOneByToken(token string) (*entity.OauthRefreshToken, *response.Errors)
	FindOneByOauthAccessTokenID(oauthAccessTokenID int) (*entity.OauthRefreshToken, *response.Errors)
}

type oauthRefreshTokenRepository struct {
	db *gorm.DB
}

func (repository *oauthRefreshTokenRepository) Create(entity entity.OauthRefreshToken) (*entity.OauthRefreshToken, *response.Errors) {
	if err := repository.db.Create(&entity).Error; err != nil {
		return nil, &response.Errors{
			Code: 500,
			Err:  nil,
		}
	}
	return &entity, nil
}

func (repository *oauthRefreshTokenRepository) Delete(entity entity.OauthRefreshToken) *response.Errors {
	//TODO implement me
	panic("implement me")
}

func (repository *oauthRefreshTokenRepository) FindOneByToken(token string) (*entity.OauthRefreshToken, *response.Errors) {
	//TODO implement me
	panic("implement me")
}

func (repository *oauthRefreshTokenRepository) FindOneByOauthAccessTokenID(oauthAccessTokenID int) (*entity.OauthRefreshToken, *response.Errors) {
	//TODO implement me
	panic("implement me")
}

func NewOauthRefreshTokenRepository(db *gorm.DB) OauthRefreshTokenRepository {
	return &oauthRefreshTokenRepository{db: db}
}
