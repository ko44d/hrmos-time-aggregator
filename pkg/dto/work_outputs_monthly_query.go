package dto

type WorkOutputsMonthlyQuery struct {
	Token   string
	Monthly string
	UserID  int
	Limit   int
	Page    int
	From    string
	To      string
}

func NewWorkOutputsMonthlyQuery(token, monthly string, userID, limit, page int, from, to string) WorkOutputsMonthlyQuery {
	return WorkOutputsMonthlyQuery{
		Token:   token,
		Monthly: monthly,
		UserID:  userID,
		Limit:   limit,
		Page:    page,
		From:    from,
		To:      to,
	}
}
