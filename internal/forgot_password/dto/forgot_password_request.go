package dto

type ForgotPasswordRequestBody struct {
	Email string `json:"email" binding:"email"`
}
