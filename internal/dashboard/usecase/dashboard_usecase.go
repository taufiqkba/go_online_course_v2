package usecase

import (
	"go_online_course_v2/internal/admin/usecase"
	"go_online_course_v2/internal/dashboard/dto"
	usecase2 "go_online_course_v2/internal/order/usecase"
	usecase3 "go_online_course_v2/internal/product/usecase"
	usecase4 "go_online_course_v2/internal/user/usecase"
)

type DashboardUseCase interface {
	GetDashboard() dto.DashboardResponseBody
}

type dashboardUseCase struct {
	adminUseCase   usecase.AdminUseCase
	orderUseCase   usecase2.OrderUseCase
	productUseCase usecase3.ProductUseCase
	userUseCase    usecase4.UserUseCase
}

func (useCase *dashboardUseCase) GetDashboard() dto.DashboardResponseBody {
	var dataDashboard dto.DashboardResponseBody

	dataDashboard.TotalAdmin = useCase.adminUseCase.TotalCountAdmin()
	dataDashboard.TotalOrder = useCase.orderUseCase.TotalCountOrder()
	dataDashboard.TotalUser = useCase.userUseCase.TotalCountUser()
	dataDashboard.TotalProduct = useCase.productUseCase.TotalCountProduct()

	return dataDashboard
}

func NewDashboardUseCase(
	adminUseCase usecase.AdminUseCase,
	orderUseCase usecase2.OrderUseCase,
	productUseCase usecase3.ProductUseCase,
	userUseCase usecase4.UserUseCase,
) DashboardUseCase {
	return &dashboardUseCase{
		adminUseCase:   adminUseCase,
		orderUseCase:   orderUseCase,
		productUseCase: productUseCase,
		userUseCase:    userUseCase,
	}
}
