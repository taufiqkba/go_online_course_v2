package dto

import "time"

type DiscountRequestBody struct {
	Name      string     `json:"name" binding:"required"`
	Code      string     `json:"code" binding:"required"`
	Quantity  int64      `json:"quantity" binding:"required"`
	Type      string     `json:"type" binding:"required"`
	Value     int64      `json:"value" binding:"required"`
	StartDate *time.Time `json:"start_date"`
	EndDate   *time.Time `json:"end_date"`
	CreatedBy *int64
	UpdatedBy *int64
}
