package dto

type VerificationEmailRequestBody struct {
	CodeVerified string `json:"code_verified" binding:"required"`
}
