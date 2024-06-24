package usecase

import (
	"github.com/ko44d/hrmos-time-aggregator/pkg/hrmos/work_outputs_monthly"
	"github.com/ko44d/hrmos-time-aggregator/pkg/repository"
)

type WorkOutputsMonthlyUsecase interface {
	Get(token string, monthly string, userId int, limit int, page int, from string, to string) ([]repository.DailyWorkData, error)
}

type workOutputsMonthlyUsecase struct {
	womr repository.WorkOutputsMonthlyRepository
}

func NewWorkOutputsMonthlyUsecase(womr repository.WorkOutputsMonthlyRepository) WorkOutputsMonthlyUsecase {
	return &workOutputsMonthlyUsecase{womr: womr}
}

func (womu *workOutputsMonthlyUsecase) Get(token string, monthly string, userId int, limit int, page int, from string, to string) ([]repository.DailyWorkData, error) {

	query := work_outputs_monthly.NewQueryParams(userId, limit, page, from, to)
	params := womu.womr.GetRequestParams(token, monthly, query)
	return womu.womr.Get(params)
}
