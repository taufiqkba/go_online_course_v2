package middleware

import (
	"github.com/gin-gonic/gin"
	"go_online_course_v2/pkg/response"
	"go_online_course_v2/pkg/utils"
	"net/http"
)

func AuthAdmin(ctx *gin.Context) {
	admin := utils.GetCurrentUser(ctx)

	if !admin.IsAdmin {
		ctx.JSON(http.StatusUnauthorized, response.Response(
			http.StatusUnauthorized,
			http.StatusText(http.StatusUnauthorized),
			"Unauthorized",
		))
		ctx.Abort()
		return
	}
	ctx.Next()
}
