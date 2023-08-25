package http

import (
	"github.com/gin-gonic/gin"
	"go_online_course_v2/internal/middleware"
	"go_online_course_v2/internal/user/dto"
	"go_online_course_v2/internal/user/usecase"
	"go_online_course_v2/pkg/response"
	"go_online_course_v2/pkg/utils"
	"net/http"
	"strconv"
)

type UserHandler struct {
	useCase usecase.UserUseCase
}

func NewUserHandler(useCase usecase.UserUseCase) *UserHandler {
	return &UserHandler{useCase: useCase}
}

func (handler *UserHandler) Route(r *gin.RouterGroup) {
	userRoute := r.Group("/api/v1")

	userRoute.Use(middleware.AuthJwt, middleware.AuthAdmin)
	{
		userRoute.GET("/user", handler.FindAll)
		userRoute.GET("/user/:id", handler.FindOneByID)
		userRoute.POST("/user", handler.Create)
		userRoute.PATCH("/user/:id", handler.Update)
		userRoute.DELETE("/user/:id", handler.Delete)
	}
}

func (handler *UserHandler) FindAll(ctx *gin.Context) {
	offset, _ := strconv.Atoi(ctx.Query("offset"))
	limit, _ := strconv.Atoi(ctx.Query("limit"))

	data := handler.useCase.FindAll(offset, limit)
	ctx.JSON(http.StatusOK, response.Response(
		http.StatusOK,
		http.StatusText(http.StatusOK),
		data,
	))
}

func (handler *UserHandler) FindOneByID(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	data, err := handler.useCase.FindOneByID(id)
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

func (handler *UserHandler) Create(ctx *gin.Context) {
	var input dto.UserRequestBody
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

	ctx.JSON(http.StatusCreated, response.Response(
		http.StatusCreated,
		http.StatusText(http.StatusCreated),
		data,
	))
}

func (handler *UserHandler) Update(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	var input dto.UserUpdateRequestBody
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

	ctx.JSON(http.StatusOK, response.Response(http.StatusOK, http.StatusText(http.StatusOK), data))
}

func (handler *UserHandler) Delete(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	err := handler.useCase.Delete(id)
	if err != nil {
		ctx.JSON(int(err.Code), response.Response(int(err.Code), http.StatusText(int(err.Code)), err.Err.Error()))
		ctx.Abort()
		return
	}
	ctx.JSON(http.StatusCreated, response.Response(
		http.StatusCreated,
		http.StatusText(http.StatusCreated),
		"OK",
	))
}
