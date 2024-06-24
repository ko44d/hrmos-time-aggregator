package dto

type WorkOutputsMonthlyQuery struct {
	Token      string
	CompanyURL string
	Monthly    string
	UserID     int
	Limit      int
	Page       int
	From       string
	To         string
}

func NewWorkOutputsMonthlyQuery(token, companyURL, monthly string, userID, limit, page int, from, to string) WorkOutputsMonthlyQuery {
	return WorkOutputsMonthlyQuery{
		Token:      token,
		CompanyURL: companyURL,
		Monthly:    monthly,
		UserID:     userID,
		Limit:      limit,
		Page:       page,
		From:       from,
		To:         to,
	}
}
