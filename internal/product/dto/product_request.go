package dto

import "mime/multipart"

type ProductRequestBody struct {
	ProductCategoryID int64                 `form:"product_category_id" binding:"required"`
	Title             string                `form:"title" binding:"required"`
	Image             *multipart.FileHeader `form:"image"`
	Video             *multipart.FileHeader `form:"video"`
	Description       string                `form:"description:" binding:"required"`
	IsHighlighted     bool                  `form:"is_highlighted"`
	Price             int64                 `form:"price" binding:"required"`
	CreatedBy         *int64
	UpdatedBy         *int64
}
