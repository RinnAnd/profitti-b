package financials

import (
	"context"
	"profitti/internal/core/domain"
	"profitti/internal/infra/database/repository/financial"
)

type FinancialService interface {
	Create(context.Context, *domain.Financial) (string, error)
	GetUserFinancials(context.Context, string) ([]*domain.Financial, error)
}

type service struct {
	repo financial.FinancialRepo
}

func New(repo financial.FinancialRepo) FinancialService {
	return &service{
		repo: repo,
	}
}

func (srv *service) Create(ctx context.Context, financial *domain.Financial) (string, error) {
	res, err := srv.repo.InsertOne(ctx, financial)
	if err != nil {
		return "", err
	}
	return res, nil
}

func (srv *service) GetUserFinancials(ctx context.Context, id string) ([]*domain.Financial, error) {
	res, err := srv.repo.SelectUserFinancials(ctx, id)
	if err != nil {
		return nil, err
	}
	return res, nil
}
