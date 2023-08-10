package dto

import "time"

type ForgotPasswordRequestBody struct {
	Email string `json:"email" binding:"email"`
}

type ForgotPasswordUpdateRequestBody struct {
	Code     string `json:"code" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type ForgotPasswordCreateRequestBody struct {
	UserID    int64      `json:"user_id"`
	Valid     bool       `json:"valid"`
	Code      string     `json:"code"`
	ExpiredAt *time.Time `json:"expired_at"`
}

type ForgotPasswordEmailRequestBody struct {
	Subject string
	Name    string
	Email   string
	Code    string
}
