package http

import (
	"github.com/gin-gonic/gin"
	"go_online_course_v2/internal/class_room/usecase"
	"go_online_course_v2/internal/middleware"
	"go_online_course_v2/pkg/response"
	"go_online_course_v2/pkg/utils"
	"net/http"
	"strconv"
)

type ClassRoomHandler struct {
	useCase usecase.ClassRoomUseCase
}

func NewClassRoomHandler() *ClassRoomHandler {
	return &ClassRoomHandler{}
}

func (handler *ClassRoomHandler) Route(r *gin.RouterGroup) {
	classRoomRoute := r.Group("/api/v1")

	classRoomRoute.Use(middleware.AuthJwt)
	{
		classRoomRoute.GET("/class_rooms", handler.FindAllByUserID)
	}

}

func (handler *ClassRoomHandler) FindAllByUserID(ctx *gin.Context) {
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
