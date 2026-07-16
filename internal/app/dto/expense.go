package dto

import (
	"profitti/internal/core/domain"
)

type Expense struct {
	Id                 string   `json:"id"`
	FinancialId        *string  `json:"financial_id,omitempty"`
	PartnershipId      *string  `json:"partnership_id,omitempty"`
	Name               string   `json:"name"`
	Description        string   `json:"description,omitempty"`
	Amount             float64  `json:"amount"`
	Category           *string  `json:"category"`
	Expense_recurrence *string  `json:"expense_recurrence,omitempty"`
	Expiration_date    *string  `json:"expiration_date,omitempty"`
	Currency           Currency `json:"currency"`
	CreatedAt          string   `json:"created_at"`
}

func (e *CreateExpense) Domain() *domain.Expense {
	return &domain.Expense{
		FinancialId:        e.FinancialId,
		PartnershipId:      e.PartnershipId,
		Name:               e.Name,
		Description:        e.Description,
		Amount:             e.Amount,
		Category:           e.CategoryId,
		Expense_recurrence: e.Expense_recurrence,
		Expiration_date:    e.Expiration_date,
		CurrencyId:         e.CurrencyId,
	}
}

type Currency struct {
	Id       int    `json:"id" validate:"required"`
	Currency string `json:"currency,omitempty"`
}

type CreateExpense struct {
	FinancialId        *string `json:"financial_id,omitempty"`
	PartnershipId      *string `json:"partnership_id,omitempty"`
	Name               string  `json:"name"`
	Description        string  `json:"description"`
	Amount             float64 `json:"amount"`
	CategoryId         *string `json:"category_id"`
	Expense_recurrence *string `json:"expense_recurrence,omitempty"`
	Expiration_date    *string `json:"expiration_date"`
	CurrencyId         int     `json:"currency_id"`
}

type CreateExpenseRes struct {
	Msg string `json:"msg"`
}

type GetExpensesByUserRes struct {
	Expenses []*Expense `json:"expenses"`
}
