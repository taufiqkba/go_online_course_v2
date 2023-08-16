package http

import (
	"github.com/gin-gonic/gin"
	"go_online_course_v2/internal/admin/dto"
	"go_online_course_v2/internal/admin/usecase"
	"go_online_course_v2/internal/middleware"
	"go_online_course_v2/pkg/response"
	"go_online_course_v2/pkg/utils"
	"net/http"
	"strconv"
)

type AdminHandler struct {
	useCase usecase.AdminUseCase
}

func NewAdminHandler(useCase usecase.AdminUseCase) *AdminHandler {
	return &AdminHandler{useCase: useCase}
}

func (handler *AdminHandler) Route(r *gin.RouterGroup) {
	adminRouter := r.Group("/api/v1")

	adminRouter.Use(middleware.AuthJwt, middleware.AuthAdmin)
	{
		adminRouter.GET("/admin", handler.FindAll)
		adminRouter.GET("/admin/:id", handler.FindByID)
		adminRouter.POST("/admin", handler.Create)
		adminRouter.PATCH("/admin/:id", handler.Update)
		adminRouter.DELETE("/admin/:id", handler.Delete)
	}
}

func (handler *AdminHandler) Create(ctx *gin.Context) {
	var input dto.AdminRequestBody

	//response error
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
	input.CreatedBy = &user.ID

	//	create data
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

	//response success created
	ctx.JSON(http.StatusCreated, response.Response(
		http.StatusCreated,
		http.StatusText(http.StatusCreated),
		"Created success",
	))
}

func (handler *AdminHandler) Update(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	var input dto.AdminRequestBody

	//response error
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
	input.UpdatedBy = &user.ID

	//	update data
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

	//response success
	ctx.JSON(http.StatusOK, response.Response(
		http.StatusOK,
		http.StatusText(http.StatusOK),
		data,
	))
}

func (handler *AdminHandler) FindAll(ctx *gin.Context) {
	offset, _ := strconv.Atoi(ctx.Query("offset"))
	limit, _ := strconv.Atoi(ctx.Query("limit"))

	data := handler.useCase.FindAll(offset, limit)

	ctx.JSON(http.StatusOK, response.Response(
		http.StatusOK,
		http.StatusText(http.StatusOK),
		data,
	))
}

func (handler *AdminHandler) FindByID(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	data, err := handler.useCase.FindByID(id)

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

func (handler *AdminHandler) Delete(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	err := handler.useCase.Delete(id)

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
		"OK",
	))
}
