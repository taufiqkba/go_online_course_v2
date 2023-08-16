package http

import (
	"github.com/gin-gonic/gin"
	"go_online_course_v2/internal/product/dto"
	"go_online_course_v2/internal/product/usecase"
	"go_online_course_v2/pkg/response"
	"go_online_course_v2/pkg/utils"
	"net/http"
	"strconv"
)

type ProductHandler struct {
	useCase usecase.ProductUseCase
}

func NewProductHandler(useCase usecase.ProductUseCase) *ProductHandler {
	return &ProductHandler{useCase: useCase}
}

func (handler *ProductHandler) Route(r *gin.RouterGroup) {

}

func (handler *ProductHandler) FindAll(ctx *gin.Context) {
	offset, _ := strconv.Atoi(ctx.Query("offset"))
	limit, _ := strconv.Atoi(ctx.Query("limit"))

	data := handler.useCase.FindAll(offset, limit)
	ctx.JSON(http.StatusOK, response.Response(
		http.StatusOK,
		http.StatusText(http.StatusOK),
		data,
	))
}

func (handler *ProductHandler) FindByID(ctx *gin.Context) {
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

func (handler *ProductHandler) Create(ctx *gin.Context) {
	//	validate input
	var input dto.ProductRequestBody

	if err := ctx.ShouldBindJSON(&input); err != nil {
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

	ctx.JSON(http.StatusOK, response.Response(
		http.StatusOK,
		http.StatusText(http.StatusOK),
		data,
	))
}

func (handler *ProductHandler) Update(ctx *gin.Context) {
	//find by id
	id, _ := strconv.Atoi(ctx.Param("id"))

	var input dto.ProductRequestBody

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, response.Response(
			http.StatusBadRequest,
			http.StatusText(http.StatusBadRequest),
			err.Error(),
		))
		ctx.Abort()
		return
	}

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

func (handler *ProductHandler) Delete(ctx *gin.Context) {
	//find by id
	id, _ := strconv.Atoi(ctx.Param("id"))

	err := handler.useCase.Delete(id)
	if err != nil {
		ctx.JSON(http.StatusOK, response.Response(
			http.StatusOK,
			http.StatusText(http.StatusOK),
			"Success",
		))
	}
}
