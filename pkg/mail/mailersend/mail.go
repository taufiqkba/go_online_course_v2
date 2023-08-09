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

	from := mailersend.From{
		Name:  "Yteam Online Course",
		Email: "verif@yteamdigital.my.id",
	}

	recipients := []mailersend.Recipient{
		{
			Email: toEmail,
		},
	}

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
					Value: "Yteam Online Course",
				},
				{
					Var:   "VerificationCode",
					Value: data.VerificationCode,
				},
			},
		},
	}

	message := ms.Email.NewMessage()

	message.SetFrom(from)
	message.SetRecipients(recipients)
	message.SetSubject(data.Subject)
	message.SetTemplateID("neqvygmr5zdl0p7w")
	message.SetSubstitutions(variables)

	res, _ := ms.Email.Send(ctx, message)

	fmt.Printf(res.Header.Get("X-Message-Id"))
}

func NewMailUseCase() Mail {
	return &mailUseCase{}
}
