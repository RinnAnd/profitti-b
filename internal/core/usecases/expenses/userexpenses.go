package expenses

import (
	"context"
	"profitti/internal/core/domain"
	"profitti/internal/infra/service/expenses"
)

type GetByUserUseCase interface {
	GetExpensesByUser(context.Context, string) ([]*domain.Expense, error)
}

type getByUserUseCase struct {
	srv expenses.ExpenseService
}

func NewGetByUserUseCase(srv expenses.ExpenseService) GetByUserUseCase {
	return &getByUserUseCase{
		srv: srv,
	}
}

func (u *getByUserUseCase) GetExpensesByUser(ctx context.Context, id string) ([]*domain.Expense, error) {
	res, err := u.srv.GetUserExpenses(ctx, id)
	if err != nil {
		return nil, err
	}
	return res, nil
}
