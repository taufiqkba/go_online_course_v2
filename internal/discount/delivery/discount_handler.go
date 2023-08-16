package delivery

import (
	"github.com/gin-gonic/gin"
	"go_online_course_v2/internal/discount/dto"
	"go_online_course_v2/internal/discount/usecase"
	"go_online_course_v2/internal/middleware"
	"go_online_course_v2/pkg/response"
	"go_online_course_v2/pkg/utils"
	"net/http"
	"strconv"
)

type DiscountHandler struct {
	useCase usecase.DiscountUseCase
}

func NewDiscountHandler(useCase usecase.DiscountUseCase) *DiscountHandler {
	return &DiscountHandler{useCase: useCase}
}

func (handler *DiscountHandler) Route(r *gin.RouterGroup) {
	routerDiscount := r.Group("/api/v1")

	routerDiscount.Use(middleware.AuthJwt, middleware.AuthAdmin)
	{
		routerDiscount.GET("/discounts", handler.FindAll)
		routerDiscount.GET("/discount/:id", handler.FindByID)
		routerDiscount.POST("/discount", handler.Create)
		routerDiscount.PATCH("/discount/:id", handler.Update)
		routerDiscount.DELETE("/discount/:id", handler.Delete)
	}
}

func (handler *DiscountHandler) FindAll(ctx *gin.Context) {
	offset, _ := strconv.Atoi(ctx.Query("offset"))
	limit, _ := strconv.Atoi(ctx.Query("limit"))

	data := handler.useCase.FindAll(offset, limit)
	ctx.JSON(http.StatusOK, response.Response(
		http.StatusOK, http.StatusText(http.StatusOK), data,
	))
}

func (handler *DiscountHandler) FindByID(ctx *gin.Context) {
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

func (handler *DiscountHandler) Create(ctx *gin.Context) {
	var input dto.DiscountRequestBody

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, response.Response(
			http.StatusBadRequest,
			http.StatusText(http.StatusBadRequest),
			err.Error(),
		))
		ctx.Abort()
		return
	}

	//set createdBy
	admin := utils.GetCurrentUser(ctx)
	input.CreatedBy = &admin.ID

	data, err := handler.useCase.Create(input)
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
		data,
	))

}

func (handler *DiscountHandler) Update(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	var input dto.DiscountRequestBody
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, response.Response(
			http.StatusBadRequest,
			http.StatusText(http.StatusBadRequest),
			err.Error(),
		))
		ctx.Abort()
		return
	}

	admin := utils.GetCurrentUser(ctx)
	input.UpdatedBy = &admin.ID

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

	ctx.JSON(http.StatusOK, response.Response(
		http.StatusOK,
		http.StatusText(http.StatusOK),
		data,
	))
}

func (handler *DiscountHandler) Delete(ctx *gin.Context) {
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
		"Success",
	))
}
