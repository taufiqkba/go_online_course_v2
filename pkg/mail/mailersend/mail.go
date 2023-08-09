package mailersend

import (
	"context"
	"fmt"
	"github.com/mailersend/mailersend-go"
	"go_online_course_v2/internal/register/dto"
	"time"
)

type Mail interface {
	SendVerificationMailer(toEmail string, data dto.EmailVerification)
}

type mailUseCase struct {
}

var APIKey string = "mlsn.24a8bac500d89c990d6a43b2757c30fc6978172e7a5c5fb64beb3ef3331b397c"

func (useCase *mailUseCase) SendVerificationMailer(toEmail string, data dto.EmailVerification) {
	ms := mailersend.NewMailersend(APIKey)

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	//
	//html := "Hello, " + data.Name + "<br>" +
	//	"Greetings from the team, your verification code is <b> " + data.VerificationCode + "</b> you got this message through your registration on our platform."

	from := mailersend.From{
		Name:  "Yteam Online Course",
		Email: "verif@yteamdigital.my.id",
	}

	recipients := []mailersend.Recipient{
		{
			Email: data.Email,
		},
	}

	variables := []mailersend.Variables{
		{
			Email: data.Email,
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
