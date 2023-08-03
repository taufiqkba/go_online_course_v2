package entity

import (
	"gorm.io/gorm"
	"time"
)

type OauthClient struct {
	ID           int64          `json:"id"`
	ClientID     string         `json:"client_id"`
	ClientSecret string         `json:"client_secret"`
	Name         string         `json:"name"`
	Redirect     string         `json:"redirect"`
	Scope        string         `json:"scope"`
	CreatedAt    *time.Time     `json:"created_at"`
	UpdatedAt    *time.Time     `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `json:"deleted_at"`
}
