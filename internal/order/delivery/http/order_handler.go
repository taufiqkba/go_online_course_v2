package http

import (
	"github.com/gin-gonic/gin"
	"go_online_course_v2/internal/middleware"
	"go_online_course_v2/internal/order/dto"
	"go_online_course_v2/internal/order/usecase"
	"go_online_course_v2/pkg/response"
	"go_online_course_v2/pkg/utils"
	"net/http"
	"strconv"
)

type OrderHandler struct {
	useCase usecase.OrderUseCase
}

func NewOrderHandler(useCase usecase.OrderUseCase) *OrderHandler {
	return &OrderHandler{useCase: useCase}
}

func (handler *OrderHandler) Route(r *gin.RouterGroup) {
	orderRoute := r.Group("/api/v1")

	orderRoute.Use(middleware.AuthJwt)
	{
		orderRoute.POST("/orders", handler.Create)
		orderRoute.GET("/orders", handler.FindAllByUserID)
		orderRoute.GET("/orders/:id", handler.FindByID)
	}
}

func (handler *OrderHandler) FindAllByUserID(ctx *gin.Context) {
	offset, _ := strconv.Atoi(ctx.Query("offset"))
	limit, _ := strconv.Atoi(ctx.Query("limit"))
	user := utils.GetCurrentUser(ctx)

	data := handler.useCase.FindAllByUserID(int(user.ID), offset, limit)
	ctx.JSON(http.StatusOK, response.Response(
		http.StatusOK,
		http.StatusText(http.StatusOK),
		data,
	))
}

func (handler *OrderHandler) FindByID(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	//user := utils.GetCurrentUser(ctx)

	data, err := handler.useCase.FindOneByID(id)
	if err != nil {
		ctx.JSON(int(err.Code), response.Response(
			int(err.Code),
			http.StatusText(int(err.Code)),
			data,
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

func (handler *OrderHandler) Create(ctx *gin.Context) {
	var input dto.OrderRequestBody

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, response.Response(
			http.StatusBadRequest,
			http.StatusText(http.StatusBadRequest),
			err,
		))
		ctx.Abort()
		return
	}
	user := utils.GetCurrentUser(ctx)

	input.UserID = user.ID
	input.Email = user.Email

	data, err := handler.useCase.Create(input)
	if err != nil {
		ctx.JSON(int(err.Code), response.Response(
			int(err.Code),
			http.StatusText(int(err.Code)),
			data,
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
