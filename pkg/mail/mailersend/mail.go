package mailersend

import (
	"context"
	"fmt"
	"github.com/mailersend/mailersend-go"
	"go_online_course_v2/internal/register/dto"
	"os"
	"time"
)

type Mail interface {
	SendVerificationMailer(toEmail string, data dto.EmailVerification)
}

type mailUseCase struct {
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

	res, _ := ms.Email.Send(ctx, message)
	fmt.Printf(res.Header.Get("X-Message-Id"))
}

func NewMailUseCase() Mail {
	return &mailUseCase{}
}
