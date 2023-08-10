package dto

import "time"

type UserRequestBody struct {
	Name      string `json:"name" binding:"required"`
	Email     string `json:"email" binding:"required,email"`
	Password  string `json:"password" binding:"required"`
	CreatedBy *int64 `json:"created_by"`
}

type UserUpdateRequestBody struct {
	Name            string     `json:"name"`
	Email           string     `json:"email"`
	EmailVerifiedAt *time.Time `json:"email_verified_at"`
	Password        *string    `json:"password"`
	UpdatedBy       *int64     `json:"updated_by"`
}
