package dto

import "profitti/internal/core/domain"

type Expense struct {
	Id                 string   `json:"id"`
	Financial_id       *string  `json:"financial_id,omitempty"`
	Partnership_id     *string  `json:"partnership_id,omitempty"`
	Name               string   `json:"name" validate:"required"`
	Description        *string  `json:"description,omitempty"`
	Amount             float64  `json:"amount" validate:"required"`
	Expense_recurrence *string  `json:"expense_recurrence,omitempty"`
	Expiration_date    *string  `json:"expiration_date,omitempty"`
	Currency_id        string   `json:"currency_id" validate:"required"`
	Currency           *string  `json:"currency,omitempty"`
	SharedPercentage   *float64 `json:"sharedPercentage,omitempty"`
	SharedCurrency     *string  `json:"sharedCurrency,omitempty"`
}

func (f *Expense) Domain() *domain.Expense {
	return &domain.Expense{
		Id:                 f.Id,
		Financial_id:       f.Financial_id,
		Partnership_id:     f.Partnership_id,
		Name:               f.Name,
		Description:        f.Description,
		Amount:             f.Amount,
		Expense_recurrence: f.Expense_recurrence,
		Expiration_date:    f.Expiration_date,
		Currency_id:        f.Currency_id,
		SharedPercentage:   f.SharedPercentage,
	}
}

type CreateExpense struct {
	Financial_id       *string `json:"financial_id"`
	Partnership_id     *string `json:"partnership_id"`
	Name               string  `json:"name"`
	Description        *string `json:"description"`
	Amount             int     `json:"amount"`
	Expense_recurrence *string `json:"expense_recurrence"`
	Expiration_date    *string `json:"expiration_date"`
	Currency_id        string  `json:"currency_id"`
}

type CreateExpenseRes struct {
	Msg string `json:"msg"`
}

type GetExpensesByUserRes struct {
	Expenses []*domain.Expense `json:"expenses"`
}
