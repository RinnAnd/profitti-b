package dto

import "profitti/internal/core/domain"

type Financial struct {
	Id         string `json:"id"`
	UserId     string `json:"user_id"`
	CurrencyId string `json:"currency_id"`
}

func (f *Financial) Domain() *domain.Financial {
	return &domain.Financial{
		Id:         f.Id,
		UserId:     f.UserId,
		CurrencyId: f.CurrencyId,
	}
}

type FinancialRes struct {
	Msg string
}

type Create struct {
	UserId     string `json:"user_id"`
	CurrencyId string `json:"currency_id"`
}
