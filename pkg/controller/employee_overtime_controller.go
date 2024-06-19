package controller

import (
	"github.com/ko44d/hrmos-time-aggregator/pkg/usecase"
	"log"
)

type EmployeeOvertimeController interface {
	Aggregate() error
}

type employeeOvertimeController struct {
	atu  usecase.AuthenticationTokenUsecase
	womu usecase.WorkOutputsMonthlyUsecase
}

func NewEmployeeOvertimeController(atu usecase.AuthenticationTokenUsecase, womu usecase.WorkOutputsMonthlyUsecase) EmployeeOvertimeController {
	return &employeeOvertimeController{atu: atu, womu: womu}
}

func (eoc *employeeOvertimeController) Aggregate() error {
	res, err := eoc.atu.Get()
	if err != nil {
		return err
	}
	log.Printf("%v\n", res)
	return nil
}
