package dto

type EmailVerification struct {
	Subject          string
	Email            string
	Name             string
	VerificationCode string
}
