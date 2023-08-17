package dto

type PaymentRequestBody struct {
	ExternalID  string `json:"external_id"`
	Amount      int64  `json:"amount"`
	PayerEmail  string `json:"payer_email"`
	Description string `json:"description"`
}
