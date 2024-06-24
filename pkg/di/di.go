package di

import (
	"github.com/ko44d/hrmos-time-aggregator/pkg/controller"
	"github.com/ko44d/hrmos-time-aggregator/pkg/repository"
	"github.com/ko44d/hrmos-time-aggregator/pkg/usecase"
	"net/http"
)

type DI struct {
}

func NewDI() *DI {
	return &DI{}
}

func (di *DI) HTTP() *http.Client {
	return http.DefaultClient
}

func (di *DI) AuthenticationTokenRepository() repository.AuthenticationTokenRepository {
	return repository.NewAuthenticationTokenRepository(di.HTTP())
}

func (di *DI) AuthenticationTokenUsecase() usecase.AuthenticationTokenUsecase {
	return usecase.NewAuthenticationTokenUsecase(di.AuthenticationTokenRepository())
}

func (di *DI) WorkOutputsMonthlyRepository() repository.WorkOutputsMonthlyRepository {
	return repository.NewWorkOutputsMonthlyRepository(di.HTTP())
}

func (di *DI) WorkOutputsMonthlyUsecase() usecase.WorkOutputsMonthlyUsecase {
	return usecase.NewWorkOutputsMonthlyUsecase(di.WorkOutputsMonthlyRepository())
}

func (di *DI) EmployeeOvertimeController() controller.EmployeeOvertimeController {
	return controller.NewEmployeeOvertimeController(di.WorkOutputsMonthlyUsecase())
}

func (di *DI) TopPageController() controller.TopPageController {
	return controller.NewTopPageController(di.AuthenticationTokenUsecase())
}
