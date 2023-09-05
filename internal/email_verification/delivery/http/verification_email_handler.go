package http

import (
	"github.com/gin-gonic/gin"
	"go_online_course_v2/internal/email_verification/dto"
	"go_online_course_v2/internal/email_verification/usecase"
	"go_online_course_v2/pkg/response"
	"net/http"
)

type VerificationEmailHandler struct {
	useCase usecase.VerificationEmailUseCase
}

func NewVerificationEmailHandler(useCase usecase.VerificationEmailUseCase) *VerificationEmailHandler {
	return &VerificationEmailHandler{useCase}
}

func (handler *VerificationEmailHandler) Route(r *gin.RouterGroup) {
	verificationEmailRoute := r.Group("/api/v1")
	verificationEmailRoute.POST("/verification_email", handler.VerificationEmail)
}

func (handler *VerificationEmailHandler) VerificationEmail(ctx *gin.Context) {
	var input dto.VerificationEmailRequestBody

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, response.Response(
			http.StatusBadRequest,
			http.StatusText(http.StatusBadRequest),
			err.Error(),
		))
		ctx.Abort()
		return
	}

	err := handler.useCase.VerificationCode(input)
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
		"OK",
	))

}
