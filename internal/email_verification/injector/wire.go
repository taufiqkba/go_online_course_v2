//go:build wireinject
// +build wireinject

package injector

import (
	"gorm.io/gorm"
)

func InitializedService(db *gorm.DB) VerificationEmail {
	wire.Build()
	return &Verificationemailhandler
}
