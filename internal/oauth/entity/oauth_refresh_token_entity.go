package entity

import (
	"gorm.io/gorm"
	"time"
)

type OauthRefreshToken struct {
	ID                 int64             `json:"id"`
	OauthAccessToken   *OauthAccessToken `gorm:"foreignKey:OauthAccessTokenID;references:ID"`
	OauthAccessTokenID *int64            `json:"oauth_access_token_id"`
	UserID             int64             `json:"user_id"`
	Token              string            `json:"token"`
	ExpiredAt          *time.Time        `json:"expired_at"`
	CreatedAt          *time.Time        `json:"created_at"`
	UpdatedAt          *time.Time        `json:"updated_at"`
	DeletedAt          gorm.DeletedAt    `json:"deleted_at"`
}
