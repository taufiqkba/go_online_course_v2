package usecase

import (
	"errors"
	"github.com/google/uuid"
	"go_online_course_v2/internal/cart/usecase"
	entity3 "go_online_course_v2/internal/discount/entity"
	usecase2 "go_online_course_v2/internal/discount/usecase"
	"go_online_course_v2/internal/order/dto"
	entity2 "go_online_course_v2/internal/order/entity"
	"go_online_course_v2/internal/order/repository"
	entity4 "go_online_course_v2/internal/order_detail/entity"
	usecase3 "go_online_course_v2/internal/order_detail/usecase"
	dto2 "go_online_course_v2/internal/payment/dto"
	usecase4 "go_online_course_v2/internal/payment/usecase"
	"go_online_course_v2/internal/product/entity"
	usecase5 "go_online_course_v2/internal/product/usecase"
	"go_online_course_v2/pkg/response"
	"strconv"
	"time"
)

type OrderUseCase interface {
	FindAllByUserID(userID int, offset int, limit int) []entity2.Order
	FindOneByID(id int) (*entity2.Order, *response.Errors)
	FindOneByExternalID(externalID string) (*entity2.Order, *response.Errors)
	Create(dto dto.OrderRequestBody) (*entity2.Order, *response.Errors)
	Update(id int, dto dto.OrderRequestBody) (*entity2.Order, *response.Errors)
	TotalCountOrder() int64
}

type orderUseCase struct {
	repository         repository.OrderRepository
	cartUseCase        usecase.CartUseCase
	discountUseCase    usecase2.DiscountUseCase
	orderDetailUseCase usecase3.OrderDetailUseCase
	paymentUseCase     usecase4.PaymentUseCase
	productUseCase     usecase5.ProductUseCase
}

func (useCase *orderUseCase) FindAllByUserID(userID int, offset int, limit int) []entity2.Order {
	return useCase.repository.FindAllByUserID(userID, offset, limit)
}

func (useCase *orderUseCase) FindOneByID(id int) (*entity2.Order, *response.Errors) {
	return useCase.repository.FindOneByID(id)
}

func (useCase *orderUseCase) FindOneByExternalID(externalID string) (*entity2.Order, *response.Errors) {
	return useCase.repository.FindOneByExternalID(externalID)
}

func (useCase *orderUseCase) Create(dto dto.OrderRequestBody) (*entity2.Order, *response.Errors) {

	//	set price
	price := 0
	totalPrice := 0
	description := ""

	var products []entity.Product
	order := entity2.Order{
		UserID: &dto.UserID,
		Status: "pending",
	}

	var dataDiscount *entity3.Discount

	//	find cart by user_id
	carts := useCase.cartUseCase.FindAllByUserID(int(dto.UserID), 1, 999)

	if len(carts) == 0 {
		//	if carts nil, check product id send by valid user or not
		if dto.ProductID != nil {
			return nil, &response.Errors{
				Code: 400,
				Err:  errors.New("carts is empty"),
			}
		}
	}

	//	validate discount
	if dto.DiscountCode != nil {
		discount, err := useCase.discountUseCase.FindByCode(*dto.DiscountCode)
		if err != nil {
			return nil, &response.Errors{
				Code: 400,
				Err:  errors.New("discount invalid"),
			}
		} else if discount.RemainingQuantity == 0 {
			if err != nil {
				return nil, &response.Errors{
					Code: 400,
					Err:  errors.New("discount invalid"),
				}
			}
		} else if discount.StartDate != nil && discount.EndDate != nil {
			if discount.StartDate.After(time.Now()) || discount.EndDate.Before(time.Now()) {
				return nil, &response.Errors{
					Code: 400,
					Err:  errors.New("discount invalid"),
				}
			}
		} else if discount.StartDate != nil {
			if discount.StartDate.After(time.Now()) {
				return nil, &response.Errors{
					Code: 400,
					Err:  errors.New("discount invalid"),
				}
			}
		} else if discount.EndDate != nil {
			if discount.StartDate.After(time.Now()) {
				return nil, &response.Errors{
					Code: 400,
					Err:  errors.New("discount invalid"),
				}
			}
		}
		//assign discount
		dataDiscount = discount
	}

	//	check carts
	if len(carts) > 0 {
		for _, carts := range carts {
			product, err := useCase.productUseCase.FindByID(int(*carts.ProductID))
			if err != nil {
				return nil, err
			}

			products = append(products, *product)
		}
	} else if dto.ProductID != nil {
		product, err := useCase.productUseCase.FindByID(int(*dto.ProductID))
		if err != nil {
			return nil, err
		}
		products = append(products, *product)
	}

	//	calculate price and send description for xendit
	for index, product := range products {
		price += int(product.Price)

		i := strconv.Itoa(index + 1)
		description += i + ". Product : " + product.Title + "<br/>"
	}
	totalPrice = price

	//	check discount data available or not
	if dataDiscount != nil {
		//	count discount logic
		if dataDiscount.Type == "rebate" {
			totalPrice = price - int(dataDiscount.Value)
		} else if dataDiscount.Type == "percent" {
			totalPrice = price - (price / 100 * int(dataDiscount.Value))
		}
		order.DiscountID = &dataDiscount.ID
	}
	order.Price = int64(price)
	order.TotalPrice = int64(totalPrice)
	order.CreatedByID = &dto.UserID

	externalID := uuid.New().String()

	order.ExternalID = externalID

	//	insert to order table
	data, err := useCase.repository.Create(order)
	if err != nil {
		return nil, err
	}

	//	insert to order_detail table
	for _, product := range products {
		orderDetail := entity4.OrderDetail{
			OrderID:     data.ID,
			ProductID:   product.ID,
			Price:       product.Price,
			CreatedByID: order.UserID,
		}
		useCase.orderDetailUseCase.Create(orderDetail)
	}

	//	hit to payment gateway xendit
	dataPayment := dto2.PaymentRequestBody{
		ExternalID:  externalID,
		Amount:      int64(totalPrice),
		PayerEmail:  dto.Email,
		Description: description,
	}

	payment, err := useCase.paymentUseCase.Create(dataPayment)
	if err != nil {
		return nil, err
	}
	data.CheckoutLink = payment.InvoiceURL

	//	update remainingQuantity
	if dto.DiscountCode != nil {
		_, err := useCase.discountUseCase.UpdateRemainingQuantity(int(dataDiscount.ID), 1, "-")
		if err != nil {
			return nil, err
		}
	}

	//	delete carts
	err = useCase.cartUseCase.DeleteByUserID(int(dto.UserID))
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (useCase *orderUseCase) Update(id int, dto dto.OrderRequestBody) (*entity2.Order, *response.Errors) {
	order, err := useCase.repository.FindOneByID(id)
	if err != nil {
		return nil, err
	}

	order.Status = dto.Status

	updateOrder, err := useCase.repository.Update(*order)

	if err != nil {
		return nil, err
	}
	return updateOrder, nil
}

func (useCase *orderUseCase) TotalCountOrder() int64 {
	//TODO implement me
	panic("implement me")
}

func NewOrderUseCase(
	repository repository.OrderRepository,
	cartUseCase usecase.CartUseCase,
	discountUseCase usecase2.DiscountUseCase,
	orderDetailUseCase usecase3.OrderDetailUseCase,
	paymentUseCase usecase4.PaymentUseCase,
	productUseCase usecase5.ProductUseCase,
) OrderUseCase {
	return &orderUseCase{
		repository:         repository,
		cartUseCase:        cartUseCase,
		discountUseCase:    discountUseCase,
		orderDetailUseCase: orderDetailUseCase,
		paymentUseCase:     paymentUseCase,
		productUseCase:     productUseCase,
	}
}
