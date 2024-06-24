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
	womu usecase.WorkOutputsMonthlyUsecase
}

func NewEmployeeOvertimeController(womu usecase.WorkOutputsMonthlyUsecase) EmployeeOvertimeController {
	return &employeeOvertimeController{womu: womu}
}

func (eoc *employeeOvertimeController) Aggregate(ctx *gin.Context) {
	token, err := ctx.Cookie("token")
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Token not found"})
		return
	}

	companyURL, err := ctx.Cookie("company_url")
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Company URL not found"})
		return
	}

	query := dto.NewWorkOutputsMonthlyQuery(token, companyURL, "2024-05", 5, 31, 0, "", "")

	data, err := eoc.womu.Get(query)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.HTML(http.StatusOK, "work_outputs.html", data)
}
