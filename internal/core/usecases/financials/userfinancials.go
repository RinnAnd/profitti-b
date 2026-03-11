package financials

import (
	"context"
	"profitti/internal/core/domain"
	"profitti/internal/infra/service/financials"
)

type GetUserFinancialsUseCase interface {
	GetFinancialsByUser(context.Context, string) ([]*domain.Financial, error)
}

type getUserFinancials struct {
	srv financials.FinancialService
}

func NewGetUserFinancials(srv financials.FinancialService) GetUserFinancialsUseCase {
	return &getUserFinancials{
		srv: srv,
	}
}

func (u *getUserFinancials) GetFinancialsByUser(ctx context.Context, id string) ([]*domain.Financial, error) {
	res, err := u.srv.GetUserFinancials(ctx, id)
	if err != nil {
		return nil, err
	}
	return res, nil
}
