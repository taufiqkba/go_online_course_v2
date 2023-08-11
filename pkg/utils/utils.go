package utils

import (
	"gorm.io/gorm"
	"math/rand"
)

func RandString(length int) string {
	var letterRune = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")

	b := make([]rune, length)
	for i := range b {
		b[i] = letterRune[rand.Intn(len(letterRune))]
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

		//	if page value <= 0 default page 1
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
