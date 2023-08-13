package http

import (
	"github.com/gin-gonic/gin"
	"go_online_course_v2/internal/oauth/dto"
	"go_online_course_v2/internal/oauth/usecase"
	"go_online_course_v2/pkg/response"
	"net/http"
)

type OauthHandler struct {
	useCase usecase.OauthUseCase
}

func NewOauthHandler(useCase usecase.OauthUseCase) *OauthHandler {
	return &OauthHandler{useCase: useCase}
}

func (handler *OauthHandler) Route(r *gin.RouterGroup) {
	oauthRouter := r.Group("/api/v1")

	oauthRouter.POST("/oauth", handler.Login)
	oauthRouter.POST("/oauth/refresh_token", handler.Refresh)
}

func (handler *OauthHandler) Login(ctx *gin.Context) {
	var input dto.LoginRequest

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, response.Response(
			http.StatusBadRequest,
			http.StatusText(http.StatusBadRequest),
			err.Error(),
		))
		ctx.Abort()
		return
	}

	//	if success, call login function
	data, err := handler.useCase.Login(input)
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

func (handler *OauthHandler) Refresh(ctx *gin.Context) {
	var input dto.RefreshTokenRequestBody

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, response.Response(
			http.StatusBadRequest,
			http.StatusText(http.StatusBadRequest),
			err.Error(),
		))
		ctx.Abort()
		return
	}

	data, err := handler.useCase.Refresh(input)

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
		http.StatusOK,
		http.StatusText(http.StatusOK),
		data,
	))
}
