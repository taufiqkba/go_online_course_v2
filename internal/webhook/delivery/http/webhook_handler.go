package http

import (
	"github.com/gin-gonic/gin"
	"go_online_course_v2/internal/webhook/dto"
	"go_online_course_v2/internal/webhook/usecase"
	"go_online_course_v2/pkg/response"
	"net/http"
)

type WebHookHandler struct {
	useCase usecase.WebHookUseCase
}

func NewWebHookHandler(useCase usecase.WebHookUseCase) *WebHookHandler {
	return &WebHookHandler{useCase: useCase}
}

func (handler *WebHookHandler) Route(r *gin.RouterGroup) {
	webHookRoute := r.Group("/api/v1")

	webHookRoute.POST("/webhook/xendit", handler.Xendit)
}

func (handler *WebHookHandler) Xendit(ctx *gin.Context) {
	var input dto.WebHookRequestBody

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, response.Response(
			http.StatusBadRequest,
			http.StatusText(http.StatusBadRequest),
			err.Error(),
		))
		ctx.Abort()
		return
	}

	err := handler.useCase.UpdatePayment(input.ID)
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
