package dto

import "profitti/internal/core/domain"

type Financial struct {
	Id         string `json:"id"`
	UserId     string `json:"user_id"`
	CurrencyId int    `json:"currency_id"`
}

func (f *CreateFinancial) Domain() *domain.Financial {
	return &domain.Financial{
		CurrencyId: f.CurrencyId,
	}
}

type FinancialRes struct {
	User       string      `json:"user"`
	Financials []Financial `json:"financials"`
}

type CreateFinancial struct {
	CurrencyId int `json:"currency_id"`
}
