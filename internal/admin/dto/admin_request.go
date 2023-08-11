package dto

type AdminRequestBody struct {
	Name      string  `json:"name" binding:"required"`
	Email     string  `json:"email" binding:"email"`
	Password  *string `json:"password"`
	CreatedBy *int64  `json:"created_by"`
	UpdatedBy *int64  `json:"updated_by"`
}
