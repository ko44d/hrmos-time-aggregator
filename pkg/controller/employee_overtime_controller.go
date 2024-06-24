package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/ko44d/hrmos-time-aggregator/pkg/dto"
	"github.com/ko44d/hrmos-time-aggregator/pkg/usecase"
	"net/http"
)

type EmployeeOvertimeController interface {
	Aggregate(ctx *gin.Context)
}

type employeeOvertimeController struct {
	atu  usecase.AuthenticationTokenUsecase
	womu usecase.WorkOutputsMonthlyUsecase
}

func NewEmployeeOvertimeController(atu usecase.AuthenticationTokenUsecase, womu usecase.WorkOutputsMonthlyUsecase) EmployeeOvertimeController {
	return &employeeOvertimeController{atu: atu, womu: womu}
}

func (eoc *employeeOvertimeController) Aggregate(ctx *gin.Context) {
	res, err := eoc.atu.Get()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	query := dto.NewWorkOutputsMonthlyQuery(res.Token, "2024-05", 0, 0, 0, "", "")

	data, err := eoc.womu.Get(query)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.HTML(http.StatusOK, "work_outputs.html", data)
}
