package http

import (
	"github.com/gin-gonic/gin"
	"go_online_course_v2/internal/forgot_password/dto"
	"go_online_course_v2/internal/forgot_password/usecase"
	"go_online_course_v2/pkg/response"
	"net/http"
)

type ForgotPasswordHandler struct {
	useCase usecase.ForgotPasswordUseCase
}

func NewForgotPasswordHandler(useCase usecase.ForgotPasswordUseCase) *ForgotPasswordHandler {
	return &ForgotPasswordHandler{useCase: useCase}
}

func (handler *ForgotPasswordHandler) Route(r *gin.RouterGroup) {
	forgotPasswordRouter := r.Group("/api/v1")

	forgotPasswordRouter.POST("/forgot_password", handler.Create)
	forgotPasswordRouter.PUT("/forgot_password", handler.Update)
}

func (handler *ForgotPasswordHandler) Create(ctx *gin.Context) {
	var input dto.ForgotPasswordRequestBody

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, response.Response(
			http.StatusBadRequest,
			http.StatusText(http.StatusBadRequest),
			err.Error(),
		))
		ctx.Abort()
		return
	}

	_, err := handler.useCase.Create(input)

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
		"Success, please check your email",
	))
}

func (handler *ForgotPasswordHandler) Update(ctx *gin.Context) {
	var input dto.ForgotPasswordUpdateRequestBody

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, response.Response(
			http.StatusBadRequest,
			http.StatusText(http.StatusBadRequest),
			err.Error(),
		))
		ctx.Abort()
		return
	}

	_, err := handler.useCase.Update(input)

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
		"Success change your password",
	))
}
