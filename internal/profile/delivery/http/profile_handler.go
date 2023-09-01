package http

import (
	"github.com/gin-gonic/gin"
	"go_online_course_v2/internal/middleware"
	dto2 "go_online_course_v2/internal/profile/dto"
	"go_online_course_v2/internal/profile/usecase"
	"go_online_course_v2/internal/user/dto"
	"go_online_course_v2/pkg/response"
	"go_online_course_v2/pkg/utils"
	"net/http"
	"strings"
)

type ProfileHandler struct {
	useCase usecase.ProfileUseCase
}

func NewProfileHandler(useCase usecase.ProfileUseCase) *ProfileHandler {
	return &ProfileHandler{useCase: useCase}
}

func (handler *ProfileHandler) Route(r *gin.RouterGroup) {
	profileRoute := r.Group("/api/v1")

	profileRoute.Use(middleware.AuthJwt)
	{
		profileRoute.GET("/profile", handler.FindProfile)
		profileRoute.PATCH("/profile", handler.Update)
		profileRoute.DELETE("/profile", handler.Deactivated)
		profileRoute.POST("/profile", handler.Logout)
	}
}

func (handler *ProfileHandler) FindProfile(ctx *gin.Context) {
	//get current user
	user := utils.GetCurrentUser(ctx)

	data, err := handler.useCase.FindProfile(int(user.ID))
	if err != nil {
		ctx.JSON(err.Code, response.Response(
			err.Code,
			http.StatusText(err.Code),
			err.Err,
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

func (handler *ProfileHandler) Update(ctx *gin.Context) {

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

	user := utils.GetCurrentUser(ctx)

	data, err := handler.useCase.Update(int(user.ID), input)
	if err != nil {
		ctx.JSON(err.Code, response.Response(
			err.Code,
			http.StatusText(err.Code),
			err.Err.Error(),
		))
	}

	ctx.JSON(http.StatusOK, response.Response(
		http.StatusOK,
		http.StatusText(http.StatusOK),
		data,
	))
}

func (handler *ProfileHandler) Deactivated(ctx *gin.Context) {
	user := utils.GetCurrentUser(ctx)

	err := handler.useCase.Deactivated(int(user.ID))
	if err != nil {
		ctx.JSON(err.Code, response.Response(
			err.Code,
			http.StatusText(err.Code),
			err.Err.Error(),
		))
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, response.Response(
		http.StatusOK,
		http.StatusText(http.StatusOK),
		http.StatusText(http.StatusOK),
	))

}

func (handler *ProfileHandler) Logout(ctx *gin.Context) {
	var input dto2.ProfileRequestBody

	_ = ctx.ShouldBindHeader(&input)

	reqToken := input.Authorization
	splitToken := strings.Split(reqToken, "Bearer ")

	err := handler.useCase.Logout(splitToken[1])
	if err != nil {
		ctx.JSON(err.Code, response.Response(
			err.Code,
			http.StatusText(err.Code),
			err.Err.Error(),
		))
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, response.Response(
		http.StatusOK,
		http.StatusText(http.StatusOK),
		http.StatusText(http.StatusOK),
	))
}
