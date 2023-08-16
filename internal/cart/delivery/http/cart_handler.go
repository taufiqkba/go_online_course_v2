package http

import (
	"github.com/gin-gonic/gin"
	"go_online_course_v2/internal/cart/dto"
	"go_online_course_v2/internal/cart/usecase"
	"go_online_course_v2/internal/middleware"
	"go_online_course_v2/pkg/response"
	"go_online_course_v2/pkg/utils"
	"net/http"
	"strconv"
)

type CartHandler struct {
	useCase usecase.CartUseCase
}

func NewCartHandler(useCase usecase.CartUseCase) *CartHandler {
	return &CartHandler{useCase: useCase}
}

func (handler *CartHandler) Route(r *gin.RouterGroup) {
	cartRoute := r.Group("/api/v1")

	const cart = "/cart"
	const cartWithID = "/cart/:id"
	cartRoute.Use(middleware.AuthJwt)
	{
		cartRoute.GET(cart, handler.FindAllByUserID)
		cartRoute.POST(cart, handler.Create)
		cartRoute.PATCH(cartWithID, handler.Update)
		cartRoute.DELETE(cartWithID, handler.Delete)
	}

}

func (handler *CartHandler) FindAllByUserID(ctx *gin.Context) {
	offset, _ := strconv.Atoi(ctx.Query("offset"))
	limit, _ := strconv.Atoi(ctx.Query("limit"))
	user := utils.GetCurrentUser(ctx)

	data := handler.useCase.FindAllByUserID(int(user.ID), offset, limit)
	ctx.JSON(http.StatusOK, response.Response(
		http.StatusOK, http.StatusText(http.StatusOK), data,
	))
}

func (handler *CartHandler) Create(ctx *gin.Context) {
	var input dto.CartRequestBody

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, response.Response(
			http.StatusBadRequest,
			http.StatusText(http.StatusBadRequest),
			err.Error(),
		))
		ctx.Abort()
		return
	}

	user := utils.GetCurrentUser(ctx)
	input.UserID = user.ID
	input.CreatedBy = input.UserID

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

func (handler *CartHandler) Update(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	var input dto.CartUpdateRequestBody
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, response.Response(
			http.StatusBadRequest,
			http.StatusText(http.StatusBadRequest),
			err.Error(),
		))
		ctx.Abort()
		return
	}

	//	get current user
	user := utils.GetCurrentUser(ctx)
	input.UserID = &user.ID

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

func (handler *CartHandler) Delete(ctx *gin.Context) {
	id, _ := strconv.Atoi("id")
	user := utils.GetCurrentUser(ctx)

	err := handler.useCase.Delete(id, int(user.ID))
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
		"Success",
	))

}
