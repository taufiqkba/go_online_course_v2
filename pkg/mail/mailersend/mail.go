package mailersend

import (
	"context"
	"github.com/mailersend/mailersend-go"
	dto2 "go_online_course_v2/internal/forgot_password/dto"
	"go_online_course_v2/internal/register/dto"
	"os"
	"time"
)

type Mail interface {
	SendVerificationMailer(toEmail string, data dto.EmailVerification)
	SendForgotPasswordMailer(toEmail string, data dto2.ForgotPasswordEmailRequestBody)
}

type mailUseCase struct {
}

func (useCase *mailUseCase) SendForgotPasswordMailer(toEmail string, data dto2.ForgotPasswordEmailRequestBody) {
	ms := mailersend.NewMailersend(os.Getenv("MAILERSEND_API_KEY"))

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	recipients := []mailersend.Recipient{
		{
			Email: toEmail,
		},
	}

	accountName := os.Getenv("MAILERSEND_NAME")
	variables := []mailersend.Variables{
		{
			Email: toEmail,
			Substitutions: []mailersend.Substitution{
				{
					Var:   "name",
					Value: data.Name,
				},
				{
					Var:   "account.name",
					Value: accountName,
				},
				{
					Var:   "code",
					Value: data.Code,
				},
			},
		},
	}

	message := ms.Email.NewMessage()

	message.SetRecipients(recipients)
	message.SetSubject(data.Subject)
	message.SetTemplateID(os.Getenv("MAILERSEND_TEMPLATE_FORGOT_PASSWORD_ID"))
	message.SetSubstitutions(variables)

	_, _ = ms.Email.Send(ctx, message)
}

func (useCase *mailUseCase) SendVerificationMailer(toEmail string, data dto.EmailVerification) {
	ms := mailersend.NewMailersend(os.Getenv("MAILERSEND_API_KEY"))

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	recipients := []mailersend.Recipient{
		{
			Email: toEmail,
		},
	}

	accountName := os.Getenv("MAILERSEND_NAME")
	variables := []mailersend.Variables{
		{
			Email: toEmail,
			Substitutions: []mailersend.Substitution{
				{
					Var:   "name",
					Value: data.Name,
				},
				{
					Var:   "account.name",
					Value: accountName,
				},
				{
					Var:   "VerificationCode",
					Value: data.VerificationCode,
				},
			},
		},
	}

	message := ms.Email.NewMessage()

	message.SetRecipients(recipients)
	message.SetSubject(data.Subject)
	message.SetTemplateID(os.Getenv("MAILERSEND_TEMPLATE_ID"))
	message.SetSubstitutions(variables)

	_, _ = ms.Email.Send(ctx, message)
}

func NewMailUseCase() Mail {
	return &mailUseCase{}
}
