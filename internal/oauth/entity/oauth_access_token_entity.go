package entity

import (
	"gorm.io/gorm"
	"time"
)

type OauthAccessToken struct {
	ID            int64          `json:"id"`
	OauthClient   *OauthClient   `gorm:"foreignKey:OauthClientID;references:ID"`
	OauthClientID *int64         `json:"oauth_client_id"`
	UserID        int64          `json:"user_id"`
	Token         string         `json:"token"`
	Scope         string         `json:"scope"`
	ExpiredAt     *time.Time     `json:"expired_at"`
	CreatedAt     *time.Time     `json:"created_at"`
	UpdatedAt     *time.Time     `json:"updated_at"`
	DeletedAt     gorm.DeletedAt `json:"deleted_at"`
}
