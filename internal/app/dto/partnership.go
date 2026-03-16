package dto

import "profitti/internal/core/domain"

type PartnershipRq struct {
	Users      []string `json:"users"`
	CurrencyId string   `json:"currency_id"`
}

func (p *PartnershipRq) Domain() *domain.Partnership {
	return &domain.Partnership{
		Users:      p.Users,
		CurrencyId: p.CurrencyId,
	}
}

type PartnershipRs struct {
	Id      string `json:"id"`
	Message string `json:"message"`
}

type PartnershipsRs struct {
	Partnerships []Partnership `json:"partnerships"`
}

type Partnership struct {
	Id         string   `json:"id"`
	Users      []string `json:"users"`
	CurrencyId string   `json:"currency_id"`
}
