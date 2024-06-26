package repository

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/ko44d/hrmos-time-aggregator/pkg/hrmos/work_outputs_monthly"
	"net/http"
	"net/url"
	"strconv"
)

type WorkOutputsMonthlyRepository interface {
	Get(params RequestParams) ([]DailyWorkData, error)
	GetRequestParams(token, companyUrl, monthly string, query work_outputs_monthly.QueryParams) RequestParams
}

type workOutputsMonthlyRepository struct {
	client *http.Client
}

type DailyWorkData struct {
	UserId                                  int64                  `json:"user_id"`
	Number                                  string                 `json:"number"`
	FullName                                string                 `json:"full_name"`
	Month                                   string                 `json:"month"`
	Day                                     string                 `json:"day"`
	Wday                                    string                 `json:"wday"`
	SegmentDisplayTitle                     string                 `json:"segment_display_title"`
	SegmentTitle                            string                 `json:"segment_title"`
	StartAt                                 string                 `json:"start_at"`
	StampingStartAt                         string                 `json:"stamping_start_at"`
	EndAt                                   string                 `json:"end_at"`
	StampingEndAt                           string                 `json:"stamping_end_at"`
	Break1StartAt                           string                 `json:"break_1_start_at"`
	StampingBreak1StartAt                   string                 `json:"stamping_break_1_start_at"`
	Break1EndAt                             string                 `json:"break_1_end_at"`
	StampingBreak1EndAt                     string                 `json:"stamping_break_1_end_at"`
	Break2StartAt                           string                 `json:"break_2_start_at"`
	StampingBreak2StartAt                   string                 `json:"stamping_break_2_start_at"`
	Break2EndAt                             string                 `json:"break_2_end_at"`
	StampingBreak2EndAt                     string                 `json:"stamping_break_2_end_at"`
	ProcedureOvertimeEndAt                  string                 `json:"procedure_overtime_end_at"`
	TotalBreakTime                          string                 `json:"total_break_time"`
	TimePaidHoliday                         string                 `json:"time_paid_holiday"`
	TotalOverWorkTime                       string                 `json:"total_over_work_time"`
	TotalOverWorkTime36                     string                 `json:"total_over_work_time_36"`
	TotalWorkingHours                       string                 `json:"total_working_hours"`
	ActualWorkingHours                      string                 `json:"actual_working_hours"`
	HoursInPrescribedWorkingHours           string                 `json:"hours_in_prescribed_working_hours"`
	HoutsInStatutoryWorkingHours            string                 `json:"hours_in_statutory_working_hours"`
	ExcessOfStatutoryWorkingHours           string                 `json:"excess_of_statutory_working_hours"`
	ExcessOfStatutoryWorkingHoursInHolidays string                 `json:"excess_of_statutory_working_hours_in_holidays"`
	HoursInStatutoryWorkingHoursInHolidays  string                 `json:"hours_in_statutory_working_hours_in_holidays"`
	LateNightOvertimeWorkingHours           string                 `json:"late_night_overtime_working_hours"`
	Status                                  int32                  `json:"status"`
	Expense                                 int32                  `json:"expense"`
	CreatedAt                               string                 `json:"created_at"`
	UpdatedAt                               string                 `json:"updated_at"`
	ExtraBreaks                             []ExtraWorkOutputBreak `json:"extra_breaks"`
	DaytimePrescribedWorkTime               int32                  `json:"daytime_prescribed_work_time"`
	MidnightPrescribedWorkTime              int32                  `json:"midnight_prescribed_work_time"`
	DaytimeStatutoryWorkOvertime            int32                  `json:"daytime_statutory_work_overtime"`
	MidnightStatutoryWorkOvertime           int32                  `json:"midnight_statutory_work_overtime"`
	DaytimeOutStatutoryWorkTime             int32                  `json:"daytime_out_statutory_work_time"`
	MidnightOutStatutoryWorkTime            int32                  `json:"midnight_out_statutory_work_time"`
	DaytimeOutStatutoryHolidayWorkTime      int32                  `json:"daytime_out_statutory_holiday_work_time"`
	MidnightOutStatutoryHolidayWorkTime     int32                  `json:"midnight_out_statutory_holiday_work_time"`
	DaytimeStatutoryHolidayWorkTime         int32                  `json:"daytime_statutory_holiday_work_time"`
	MidnightStatutoryHolidayWorkTime        int32                  `json:"midnight_statutory_holiday_work_time"`
	PaidHolidayWithTimePaid                 float64                `json:"paid_holiday_with_time_paid"`
	Notes                                   string                 `json:"notes"`
	ImageAttachmentIds                      []int                  `json:"image_attachment_ids"`
	PdfAttachmentIds                        []int                  `json:"pdf_attachment_ids"`
	Work                                    Work                   `json:"work"`
}

type ExtraWorkOutputBreak struct {
	StartAt         string `json:"start_at""`
	EndAt           string `json:"end_at"`
	StampingStartAt string `json:"stamping_start_at"`
	StampingEndAt   string `json:"stamping_end_at"`
}

type Work struct {
	Id                  int64  `json:"id"`
	UserId              int64  `json:"user_id"`
	Day                 string `json:"day"`
	MainOperationRatio  string `json:"main_operation_ratio"`
	SubOperationRatio   string `json:"sub_operation_ratio"`
	OtherOperationRatio string `json:"other_operation_ratio"`
}

type RequestParams struct {
	token      string
	companyURL string
	monthly    string
	query      work_outputs_monthly.QueryParams
}

func NewWorkOutputsMonthlyRepository(client *http.Client) WorkOutputsMonthlyRepository {
	return &workOutputsMonthlyRepository{client: client}
}

func (womr *workOutputsMonthlyRepository) GetRequestParams(token, companyURL, monthly string, query work_outputs_monthly.QueryParams) RequestParams {
	return RequestParams{
		token:      token,
		companyURL: companyURL,
		monthly:    monthly,
		query:      query,
	}
}

func (womr *workOutputsMonthlyRepository) Get(params RequestParams) ([]DailyWorkData, error) {
	u, err := url.Parse(fmt.Sprintf("%s://%s/api/%s/v1/work_outputs/monthly/%s", "https", "ieyasu.co", params.companyURL, params.monthly))
	if err != nil {
		return nil, err
	}

	u.RawQuery = params.query.ToValues().Encode()

	req, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Token %s", params.token))
	req.Header.Set("Content-Type", "application/json")

	res, err := womr.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, errors.New(fmt.Sprintf("failed to fetch work outputs monthly data. status code: %s", strconv.Itoa(res.StatusCode)))
	}

	var dwd []DailyWorkData
	if err := json.NewDecoder(res.Body).Decode(&dwd); err != nil {
		return nil, err
	}
	return dwd, nil
}
