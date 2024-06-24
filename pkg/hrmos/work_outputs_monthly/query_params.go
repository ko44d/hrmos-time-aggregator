package work_outputs_monthly

import (
	"errors"
	"net/url"
	"strconv"
)

type QueryParams interface {
	ToValues() url.Values
	Validate() error
}

type queryParams struct {
	UserID int
	Limit  int
	Page   int
	From   string
	To     string
}

func NewQueryParams(userID, limit, page int, from, to string) (QueryParams, error) {
	qp := &queryParams{
		UserID: userID,
		Limit:  limit,
		Page:   page,
		From:   from,
		To:     to,
	}
	if err := qp.Validate(); err != nil {
		return nil, err
	}
	return qp, nil
}

func (qp *queryParams) ToValues() url.Values {
	v := url.Values{}
	if qp.UserID != 0 {
		v.Set("user_id", strconv.Itoa(qp.UserID))
	}
	v.Set("limit", strconv.Itoa(qp.limitOrDefault()))
	v.Set("page", strconv.Itoa(qp.pageOrDefault()))
	if qp.From != "" {
		v.Set("from", qp.From)
	}
	if qp.To != "" {
		v.Set("to", qp.To)
	}
	return v
}

func (qp *queryParams) Validate() error {
	if qp.Limit < 0 {
		return errors.New("limit must be a positive integer")
	}
	if qp.Page < 0 {
		return errors.New("page must be a positive integer")
	}
	return nil
}

func (qp *queryParams) limitOrDefault() int {
	if qp.Limit != 0 {
		return qp.Limit
	}
	return 25
}

func (qp *queryParams) pageOrDefault() int {
	if qp.Page != 0 {
		return qp.Page
	}
	return 1
}
