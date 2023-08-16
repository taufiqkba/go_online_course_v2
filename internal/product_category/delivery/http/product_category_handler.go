package http

import (
	"github.com/gin-gonic/gin"
	"go_online_course_v2/internal/middleware"
	"go_online_course_v2/internal/product_category/dto"
	"go_online_course_v2/internal/product_category/usecase"
	"go_online_course_v2/pkg/response"
	"go_online_course_v2/pkg/utils"
	"net/http"
	"strconv"
)

type ProductCategoryHandler struct {
	useCase usecase.ProductCategoryUseCase
}

func NewProductCategoryHandler(useCase usecase.ProductCategoryUseCase) *ProductCategoryHandler {
	return &ProductCategoryHandler{useCase: useCase}
}

func (handler *ProductCategoryHandler) Route(r *gin.RouterGroup) {
	productCategoryRouter := r.Group("/api/v1")

	productCategoryRouter.Use(middleware.AuthJwt, middleware.AuthAdmin)
	{
		productCategoryRouter.GET("/product_categories", handler.FindAll)
		productCategoryRouter.GET("/product_category/:id", handler.FindByID)
		productCategoryRouter.POST("/product_category", handler.Create)
		productCategoryRouter.PATCH("/product_category/:id", handler.Update)
		productCategoryRouter.DELETE("/product_category/:id", handler.Delete)
	}
}

func (handler *ProductCategoryHandler) FindAll(ctx *gin.Context) {
	offset, _ := strconv.Atoi(ctx.Query("offset"))
	limit, _ := strconv.Atoi(ctx.Query("limit"))

	data := handler.useCase.FindAll(offset, limit)
	ctx.JSON(http.StatusOK, response.Response(
		http.StatusOK,
		http.StatusText(http.StatusOK),
		data,
	))
}

func (handler *ProductCategoryHandler) FindByID(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	data, err := handler.useCase.FindByID(id)
	if err != nil {
		ctx.JSON(int(err.Code), response.Response(
			int(err.Code),
			http.StatusText(int(err.Code)),
			err.Err.Error(),
		))
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, response.Response(
		http.StatusOK,
		http.StatusText(http.StatusOK),
		data,
	))
}

func (handler *ProductCategoryHandler) Create(ctx *gin.Context) {
	//	validate input
	var input dto.ProductCategoryRequestBody

	if err := ctx.ShouldBind(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, response.Response(
			http.StatusBadRequest,
			http.StatusText(http.StatusBadRequest),
			err.Error(),
		))
		ctx.Abort()
		return
	}

	admin := utils.GetCurrentUser(ctx)
	input.CreatedBy = &admin.ID

	data, err := handler.useCase.Create(input)
	if err != nil {
		ctx.JSON(int(err.Code), response.Response(
			int(err.Code),
			http.StatusText(int(err.Code)),
			err.Err.Error(),
		))
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusCreated, response.Response(
		http.StatusCreated,
		http.StatusText(http.StatusCreated),
		data,
	))
}

func (handler *ProductCategoryHandler) Update(ctx *gin.Context) {
	//	find data by id
	id, _ := strconv.Atoi(ctx.Param("id"))

	var input dto.ProductCategoryRequestBody

	if err := ctx.ShouldBind(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, response.Response(
			http.StatusBadRequest,
			http.StatusText(http.StatusBadRequest),
			err.Error(),
		))
		ctx.Abort()
		return
	}

	admin := utils.GetCurrentUser(ctx)

	input.UpdatedBy = &admin.ID

	data, err := handler.useCase.Update(id, input)
	if err != nil {
		ctx.JSON(int(err.Code), response.Response(
			int(err.Code),
			http.StatusText(int(err.Code)),
			err.Err.Error(),
		))
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, response.Response(
		http.StatusOK,
		http.StatusText(http.StatusOK),
		data,
	))

}

func (handler *ProductCategoryHandler) Delete(ctx *gin.Context) {
	//	find data by id
	id, _ := strconv.Atoi(ctx.Param("id"))

	err := handler.useCase.Delete(id)
	if err != nil {
		ctx.JSON(int(err.Code), response.Response(
			int(err.Code),
			http.StatusText(int(err.Code)),
			err.Err.Error(),
		))
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, response.Response(
		http.StatusOK,
		http.StatusText(http.StatusOK),
		nil,
	))
}
