package domain

type Expense struct {
	Id                 string
	FinancialId        *string
	PartnershipId      *string
	Name               string
	Description        string
	Amount             float64
	Category           *string
	Expense_recurrence *string
	Expiration_date    *string
	CreatedAt          string
	CurrencyId         int
}
