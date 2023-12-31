package utils

import (
	"github.com/gin-gonic/gin"
	"go_online_course_v2/internal/oauth/dto"
	"gorm.io/gorm"
	"math/rand"
	"path/filepath"
	"time"
)

func RandString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyz" +
		"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	var seededRand *rand.Rand = rand.New(
		rand.NewSource(time.Now().UnixNano()))

	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func RandNumber(length int) string {
	var letterRune = []rune("1234567890")

	b := make([]rune, length)
	for i := range b {
		b[i] = letterRune[rand.Intn(len(letterRune))]
	}
	return string(b)
}

func Paginate(offset int, limit int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		page := offset

		//	if page value <= 0  set default page to 1
		if page <= 0 {
			page = 1
		}

		pageSize := limit
		switch {
		case pageSize > 100:
			pageSize = 100

		case pageSize <= 0:
			pageSize = 10
		}

		offset := (page - 1) * limit
		return db.Offset(offset).Limit(pageSize)
	}
}

func GetCurrentUser(ctx *gin.Context) *dto.ClaimResponse {
	user, _ := ctx.Get("user")
	return user.(*dto.ClaimResponse)
}

func GetFileName(fileName string) string {
	file := filepath.Base(fileName)
	return file[:len(file)-len(filepath.Ext(file))]
}
