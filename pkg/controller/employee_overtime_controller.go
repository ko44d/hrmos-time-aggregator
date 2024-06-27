package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ko44d/hrmos-time-aggregator/pkg/dto"
	"github.com/ko44d/hrmos-time-aggregator/pkg/repository"
	"github.com/ko44d/hrmos-time-aggregator/pkg/usecase"
	"net/http"
	"strconv"
	"strings"
	"time"
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

var userid_list = []int{
	7, 12, 14, 21, 24,
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
	monthly := ctx.Query("monthly")
	if monthly == "" {
		currentTime := time.Now()
		monthly = currentTime.Format("2006-01")
	}

	dailyWorkDatas := []struct {
		DailyWorkData []repository.DailyWorkData
		TotalOvertime string
	}{}

	var totalOvertime string
	for _, ui := range userid_list {
		query := dto.NewWorkOutputsMonthlyQuery(token, companyURL, monthly, ui, 31, 1, "", "")
		data, err := eoc.womu.Get(query)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		totalOvertimeMinutes := 0
		for _, record := range data {
			if record.SegmentTitle == "直行直帰(残業有)" || record.SegmentTitle == "直行(残業有)" || record.SegmentTitle == "直帰(残業有)" {
				parts := strings.Split(record.TotalOverWorkTime, ":")
				if len(parts) == 2 {
					hours, err := strconv.Atoi(parts[0])
					if err != nil {
						ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
					}
					minutes, err := strconv.Atoi(parts[1])
					if err != nil {
						ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
					}
					totalOvertimeMinutes += hours*60 + minutes
				}
			}
		}

		totalHours := totalOvertimeMinutes / 60
		totalMinutes := totalOvertimeMinutes % 60
		totalOvertime = strconv.Itoa(totalHours) + ":" + fmt.Sprintf("%02d", totalMinutes)

		dailyWorkData := struct {
			DailyWorkData []repository.DailyWorkData
			TotalOvertime string
		}{
			DailyWorkData: data,
			TotalOvertime: totalOvertime,
		}

		dailyWorkDatas = append(dailyWorkDatas, dailyWorkData)
	}
	ctx.HTML(http.StatusOK, "work_outputs.html", gin.H{
		"dailyWorkDatas": dailyWorkDatas,
		"monthly":        monthly,
	})
}
