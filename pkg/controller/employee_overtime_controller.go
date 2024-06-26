package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ko44d/hrmos-time-aggregator/pkg/dto"
	"github.com/ko44d/hrmos-time-aggregator/pkg/usecase"
	"net/http"
	"regexp"
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

var timeFormat = regexp.MustCompile(`^\d{1,2}:\d{2}$`)

var pattern = regexp.MustCompile(`^(直行直帰|直行|直帰)\(残業有\)$`)

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

	query := dto.NewWorkOutputsMonthlyQuery(token, companyURL, monthly, 7, 31, 1, "", "")
	data, err := eoc.womu.Get(query)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var totalMinutes int
	for _, d := range data {
		if pattern.MatchString(d.SegmentTitle) {
			if !timeFormat.MatchString(d.TotalOverWorkTime) {
				ctx.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Errorf("invalid time format")})
			}
			parts := strings.Split(d.TotalOverWorkTime, ":")
			hours, err := strconv.Atoi(parts[0])
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			}
			minutes, err := strconv.Atoi(parts[1])
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			}
			totalMinutes += hours*60 + minutes
		}
	}

	hours := totalMinutes / 60
	minutes := totalMinutes % 60

	ctx.HTML(http.StatusOK, "work_outputs.html", gin.H{
		"data":    data,
		"monthly": monthly,
		"hours":   hours,
		"minutes": minutes,
	})
}
