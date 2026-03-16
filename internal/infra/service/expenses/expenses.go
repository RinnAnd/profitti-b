package expenses

import (
	"context"
	"profitti/internal/core/domain"
	expense "profitti/internal/infra/database/repository/expenses"
)

type ExpenseService interface {
	Create(context.Context, *domain.Expense) (string, error)
	GetUserExpenses(context.Context, string) ([]*domain.Expense, error)
}

type service struct {
	repo expense.ExpenseRepo
}

func New(repo expense.ExpenseRepo) ExpenseService {
	return &service{
		repo: repo,
	}
}

func (s *service) Create(ctx context.Context, e *domain.Expense) (string, error) {
	res, err := s.repo.InsertOne(ctx, e)
	if err != nil {
		return "", err
	}
	return res, nil
}

func (s *service) GetUserExpenses(ctx context.Context, id string) ([]*domain.Expense, error) {
	res, err := s.repo.SelectUserExpenses(ctx, id)
	if err != nil {
		return nil, err
	}
	return res, nil
}
