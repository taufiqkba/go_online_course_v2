package http

import (
	"github.com/gin-gonic/gin"
	"go_online_course_v2/internal/dashboard/usecase"
	"go_online_course_v2/internal/middleware"
	"go_online_course_v2/pkg/response"
	"net/http"
)

type DashboardHandler struct {
	useCase usecase.DashboardUseCase
}

func NewDashboardHandler(useCase usecase.DashboardUseCase) *DashboardHandler {
	return &DashboardHandler{useCase: useCase}
}

func (handler *DashboardHandler) Route(r *gin.RouterGroup) {
	dashboardRoute := r.Group("/api/v1")

	dashboardRoute.Use(middleware.AuthJwt, middleware.AuthAdmin)
	{
		dashboardRoute.GET("/dashboard", handler.GetDataDashboard)
	}
}

func (handler *DashboardHandler) GetDataDashboard(ctx *gin.Context) {
	data := handler.useCase.GetDashboard()

	ctx.JSON(http.StatusOK, response.Response(http.StatusOK, http.StatusText(http.StatusOK), data))
}
