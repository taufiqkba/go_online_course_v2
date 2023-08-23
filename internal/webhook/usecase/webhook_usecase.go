package usecase

import (
	"errors"
	"fmt"
	"github.com/xendit/xendit-go"
	"github.com/xendit/xendit-go/invoice"
	"go_online_course_v2/internal/class_room/dto"
	usecase2 "go_online_course_v2/internal/class_room/usecase"
	dto2 "go_online_course_v2/internal/order/dto"
	"go_online_course_v2/internal/order/usecase"
	"go_online_course_v2/pkg/response"
	"os"
	"strings"
)

type WebHookUseCase interface {
	UpdatePayment(id string) *response.Errors
}

type webHookUseCase struct {
	orderUseCase     usecase.OrderUseCase
	classRoomUseCase usecase2.ClassRoomUseCase
}

func (useCase *webHookUseCase) UpdatePayment(id string) *response.Errors {
	//	check data from xendit
	params := invoice.GetParams{ID: id}

	dataXendit, err := invoice.Get(&params)
	if err != nil {
		return &response.Errors{
			Code: 500,
			Err:  err,
		}
	}

	if dataXendit == nil {
		return &response.Errors{
			Code: 404,
			Err:  errors.New("order not found"),
		}
	}

	dataOrder, errOrder := useCase.orderUseCase.FindOneByExternalID(dataXendit.ExternalID)
	if errOrder != nil {
		return errOrder
	}

	if dataOrder == nil {
		return &response.Errors{
			Code: 404,
			Err:  errors.New("order not found"),
		}
	}

	if dataOrder.Status == "settled" {
		return &response.Errors{
			Code: 400,
			Err:  errors.New("payment has been already processed"),
		}
	}

	if dataOrder.Status != "paid" {
		if dataXendit.Status == "PAID" || dataXendit.Status == "SETTLED" {
			//	add data to classrooms
			for _, orderDetail := range dataOrder.OrderDetails {
				dataClassRoom := dto.ClassRoomRequestBody{
					UserID:    *dataOrder.UserID,
					ProductID: *orderDetail.ProductID,
				}
				_, err := useCase.classRoomUseCase.Create(dataClassRoom)
				if err != nil {
					fmt.Println(err)
				}
			}

			//	trigger another notification to email, ect
		}
	}

	//	update order data
	order := dto2.OrderRequestBody{
		Status: strings.ToLower(dataXendit.Status),
	}

	useCase.orderUseCase.Update(int(dataOrder.ID), order)
	return nil
}

func NewWebHookUseCase(orderUseCase usecase.OrderUseCase, classRoomUseCase usecase2.ClassRoomUseCase) WebHookUseCase {
	xendit.Opt.SecretKey = os.Getenv("XENDIT_APIKEY")
	return &webHookUseCase{orderUseCase: orderUseCase, classRoomUseCase: classRoomUseCase}
}
