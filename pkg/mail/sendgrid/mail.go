package sendgrid

import (
	"bytes"
	"fmt"
	"github.com/sendgrid/sendgrid-go"
	_ "github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	"go_online_course_v2/internal/register/dto"
	"html/template"
	"os"
	"path/filepath"
)

type Mail interface {
	SendVerification(toEmail string, data dto.EmailVerification)
}

type mailUseCase struct {
}

func (usecase *mailUseCase) sendMail(toEmail string, result string, subject string) {
	from := mail.NewEmail(os.Getenv("MAIL_SENDER_NAME"), os.Getenv("MAIL_SENDER_NAME"))
	to := mail.NewEmail(toEmail, toEmail)

	messages := mail.NewSingleEmail(from, subject, to, "", result)

	client := sendgrid.NewSendClient(os.Getenv("MAIL_KEY"))
	resp, err := client.Send(messages)
	if err != nil {
		fmt.Println("err")
	} else if resp.StatusCode != 200 {
		fmt.Println(resp)
	} else {
		fmt.Printf("success send email to %s\n", toEmail)
	}
}

func (usecase *mailUseCase) SendVerification(toEmail string, data dto.EmailVerification) {
	cwd, _ := os.Getwd()
	templateFile := filepath.Join(cwd, "/templates/email/verification_email.gohtml")
	result, err := ParseTemplate(templateFile, data)
	if err != nil {
		fmt.Println(err)
	} else {
		usecase.sendMail(toEmail, result, data.Subject)
	}
}

func ParseTemplate(templateName string, data interface{}) (string, error) {
	t, err := template.ParseFiles(templateName)
	if err != nil {
		return "", err
	}

	buf := new(bytes.Buffer)

	if err = t.Execute(buf, data); err != nil {
		return "", err
	}
	return buf.String(), nil
}

func NewMailUseCase() Mail {
	return &mailUseCase{}
}
