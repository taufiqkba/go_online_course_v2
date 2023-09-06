package http

import (
	"github.com/gin-gonic/gin"
	"go_online_course_v2/internal/verification_email/dto"
	"go_online_course_v2/internal/verification_email/usecase"
	"go_online_course_v2/pkg/response"
	"net/http"
)

type VerificationEmailHandler struct {
	usecase usecase.VerificationEmailUseCase
}

func NewVerificationEmailHandler(usecase usecase.VerificationEmailUseCase) *VerificationEmailHandler {
	return &VerificationEmailHandler{usecase}
}

func (handler *VerificationEmailHandler) Route(r *gin.RouterGroup) {
	verificationEmailRouter := r.Group("/api/v1")

	verificationEmailRouter.POST("/verification_emails", handler.VerificationEmail)
}

func (handler *VerificationEmailHandler) VerificationEmail(ctx *gin.Context) {
	// Validate input
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

	err := handler.usecase.VerificationCode(input)

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
		http.StatusText(http.StatusOK),
	))
}
