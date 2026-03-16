package domain

type Expense struct {
	Id                 string
	Financial_id       *string
	Partnership_id     *string
	Name               string
	Description        *string
	Amount             float64
	Expense_recurrence *string
	Expiration_date    *string
	Currency_id        string
	Currency           string
	SharedPercentage   *float64
	SharedCurrency     *string
}
