package repository

import (
	"go_online_course_v2/internal/oauth/entity"
	"go_online_course_v2/pkg/response"
	"gorm.io/gorm"
)

type OauthClientRepository interface {
	FindByClientIDAndClientSecret(clientID string, clientSecret string) (*entity.OauthClient, *response.Errors)
}

type oauthClientRepository struct {
	db *gorm.DB
}

// FindByClientIDAndClientSecret implements OauthClientRepository
func (repository *oauthClientRepository) FindByClientIDAndClientSecret(clientID string, clientSecret string) (*entity.OauthClient, *response.Errors) {
	var oauthClient entity.OauthClient

	if err := repository.db.
		Where("client_id = ?", clientID).
		Where("client_secret = ?", clientSecret).
		First(&oauthClient).
		Error; err != nil {
		return nil, &response.Errors{
			Code: 500,
			Err:  err,
		}
	}
	return &oauthClient, nil
}

func NewOauthClientRepository(db *gorm.DB) OauthClientRepository {
	return &oauthClientRepository{db: db}
}
