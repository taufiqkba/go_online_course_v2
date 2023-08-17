package dto

type OrderRequestBody struct {
	DiscountCode *string `json:"discount_code"`
	ProductID    *int64  `json:"product_id"`
	UserID       int64   `json:"user_id"`
	Email        string  `json:"email"`
	Status       string  `json:"status"`
}
