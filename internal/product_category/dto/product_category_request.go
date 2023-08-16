package dto

import "mime/multipart"

type ProductCategoryRequestBody struct {
	Name      string                `form:"name" binding:"required"`
	Image     *multipart.FileHeader `form:"image"`
	CreatedBy *int64
	UpdatedBy *int64
}
