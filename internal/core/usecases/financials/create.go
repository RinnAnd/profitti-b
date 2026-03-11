package financials

import (
	"context"
	"profitti/internal/core/domain"
	"profitti/internal/infra/service/financials"
)

type CreateUseCase interface {
	CreateFinancial(context.Context, *domain.Financial) (string, error)
}

type createFinancialUseCase struct {
	srv financials.FinancialService
}

func NewCreateUseCase(srv financials.FinancialService) CreateUseCase {
	return &createFinancialUseCase{
		srv: srv,
	}
}

func (u *createFinancialUseCase) CreateFinancial(ctx context.Context, financial *domain.Financial) (string, error) {
	res, err := u.srv.Create(ctx, financial)
	if err != nil {
		return "", err
	}
	return res, nil
}
