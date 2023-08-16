package dto

type CartRequestBody struct {
	ProductID int64 `json:"product_id,omitempty" binding:"required,number"`
	UserID    int64 `json:"user_id"`
	CreatedBy int64
	UpdatedBy int64
}

type CartUpdateRequestBody struct {
	IsChecked bool   `json:"is_checked,omitempty"`
	UserID    *int64 `json:"user_id"`
}
