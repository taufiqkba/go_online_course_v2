package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"go_online_course_v2/internal/oauth/dto"
	"go_online_course_v2/pkg/response"
	"net/http"
	"os"
	"strings"
)

type Header struct {
	Authorization string `header:"authorization" binding:"required"`
}

func AuthJwt(ctx *gin.Context) {
	var input Header

	if err := ctx.ShouldBindHeader(&input); err != nil {
		ctx.JSON(http.StatusUnauthorized, response.Response(
			http.StatusUnauthorized,
			http.StatusText(http.StatusUnauthorized),
			"Unauthorized",
		))
		ctx.Abort()
		return
	}

	//	Split token
	reqToken := input.Authorization
	splitToken := strings.Split(reqToken, "Bearer ")

	if len(splitToken) != 2 {
		ctx.JSON(http.StatusUnauthorized, response.Response(
			http.StatusUnauthorized,
			http.StatusText(http.StatusUnauthorized),
			"Unauthorized",
		))
		ctx.Abort()
		return
	}

	//[0] = Bearer
	//[1] = token
	reqToken = splitToken[1]
	claims := &dto.ClaimResponse{}

	token, err := jwt.ParseWithClaims(reqToken, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil {
		ctx.JSON(http.StatusUnauthorized, response.Response(
			http.StatusUnauthorized,
			http.StatusText(http.StatusUnauthorized),
			"Unauthorized",
		))
		ctx.Abort()
		return
	}

	if !token.Valid {
		ctx.JSON(http.StatusUnauthorized, response.Response(
			http.StatusUnauthorized,
			http.StatusText(http.StatusUnauthorized),
			"Unauthorized",
		))
		ctx.Abort()
		return
	}

	claims = token.Claims.(*dto.ClaimResponse)
	ctx.Set("user", claims)
	ctx.Next()
}
