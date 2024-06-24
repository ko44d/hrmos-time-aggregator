package work_outputs_monthly

import (
	"net/url"
	"strconv"
)

type QueryParams interface {
	ToValues() url.Values
	LimitOrDefault() int
	PageOrDefault() int
}

type queryParams struct {
	UserID int
	Limit  int
	Page   int
	From   string
	To     string
}

func NewQueryParams(userID, limit, page int, from, to string) QueryParams {
	return &queryParams{
		UserID: userID,
		Limit:  limit,
		Page:   page,
		From:   from,
		To:     to,
	}
}

func (qp *queryParams) ToValues() url.Values {
	v := url.Values{}
	if qp.UserID != 0 {
		v.Set("user_id", strconv.Itoa(qp.UserID))
	}
	v.Set("limit", strconv.Itoa(qp.LimitOrDefault()))
	v.Set("page", strconv.Itoa(qp.PageOrDefault()))
	if qp.From != "" {
		v.Set("from", qp.From)
	}
	if qp.To != "" {
		v.Set("to", qp.To)
	}
	return v
}

func (qp *queryParams) LimitOrDefault() int {
	if qp.Limit != 0 {
		return qp.Limit
	}
	return 25
}

func (qp *queryParams) PageOrDefault() int {
	if qp.Page != 0 {
		return qp.Page
	}
	return 1
}
