package usecase

import (
	"github.com/ko44d/hrmos-time-aggregator/pkg/dto"
	"github.com/ko44d/hrmos-time-aggregator/pkg/hrmos/work_outputs_monthly"
	"github.com/ko44d/hrmos-time-aggregator/pkg/repository"
)

type WorkOutputsMonthlyUsecase interface {
	Get(query dto.WorkOutputsMonthlyQuery) ([]repository.DailyWorkData, error)
}

type workOutputsMonthlyUsecase struct {
	womr repository.WorkOutputsMonthlyRepository
}

func NewWorkOutputsMonthlyUsecase(womr repository.WorkOutputsMonthlyRepository) WorkOutputsMonthlyUsecase {
	return &workOutputsMonthlyUsecase{womr: womr}
}

func (womu *workOutputsMonthlyUsecase) Get(query dto.WorkOutputsMonthlyQuery) ([]repository.DailyWorkData, error) {
	q, err := work_outputs_monthly.NewQueryParams(query.UserID, query.Limit, query.Page, query.From, query.To)
	if err != nil {
		return nil, err
	}
	params := womu.womr.GetRequestParams(query.Token, query.CompanyURL, query.Monthly, q)
	return womu.womr.Get(params)
}
