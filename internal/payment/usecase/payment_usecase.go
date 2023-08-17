package usecase

import (
	"github.com/xendit/xendit-go"
	"github.com/xendit/xendit-go/invoice"
	"go_online_course_v2/internal/payment/dto"
	"go_online_course_v2/pkg/response"
	"os"
)

type PaymentUseCase interface {
	Create(dto dto.PaymentRequestBody) (*xendit.Invoice, *response.Errors)
}

type paymentUseCase struct {
}

func (useCase *paymentUseCase) Create(dto dto.PaymentRequestBody) (*xendit.Invoice, *response.Errors) {
	data := invoice.CreateParams{
		ExternalID:      dto.ExternalID,
		Amount:          float64(dto.Amount),
		Description:     dto.Description,
		PayerEmail:      dto.PayerEmail,
		ShouldSendEmail: nil,
		Customer: xendit.InvoiceCustomer{
			Email: dto.PayerEmail,
		},
		CustomerNotificationPreference: xendit.InvoiceCustomerNotificationPreference{
			InvoiceCreated:  []string{"email"},
			InvoiceReminder: []string{"email"},
			InvoicePaid:     []string{"email"},
			InvoiceExpired:  []string{"email"},
		},
		InvoiceDuration:    86400,
		SuccessRedirectURL: os.Getenv("XENDIT_SUCCESS_URL"),
	}

	resp, err := invoice.Create(&data)
	if err != nil {
		return nil, &response.Errors{
			Code: 500,
			Err:  err,
		}
	}
	return resp, nil
}

func NewPaymentUseCase() PaymentUseCase {
	//Setup xendit payment gateway
	xendit.Opt.SecretKey = os.Getenv("XENDIT_APIKEY")
	return &paymentUseCase{}
}
