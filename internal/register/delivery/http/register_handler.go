package http

import (
	"github.com/gin-gonic/gin"
	"go_online_course_v2/internal/register/usecase"
	"go_online_course_v2/internal/user/dto"
	"go_online_course_v2/pkg/response"
	"net/http"
)

type RegisterHandler struct {
	useCase usecase.RegisterUseCase
}

func NewRegisterHandler(useCase usecase.RegisterUseCase) *RegisterHandler {
	return &RegisterHandler{useCase: useCase}
}

func (handler *RegisterHandler) Route(r *gin.RouterGroup) {
	r.POST("/api/v1/register", handler.Register)
}

func (handler *RegisterHandler) Register(ctx *gin.Context) {
	// Input validation
	var registerRequestInput dto.UserRequestBody

	if err := ctx.ShouldBindJSON(&registerRequestInput); err != nil {
		ctx.JSON(http.StatusBadRequest, response.Response(
			http.StatusBadRequest,
			http.StatusText(http.StatusBadRequest),
			err.Error(),
		))
		ctx.Abort()
		return
	}

	err := handler.useCase.Register(registerRequestInput)

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
		"Success, please check your email",
	))
}
