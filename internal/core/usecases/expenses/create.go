package expenses

import (
	"context"
	"profitti/internal/core/domain"
	"profitti/internal/infra/service/expenses"
)

type CreateUseCase interface {
	CreateExpense(context.Context, *domain.Expense) (string, error)
}

type createExpenseUseCase struct {
	srv expenses.ExpenseService
}

func NewCreateUseCase(srv expenses.ExpenseService) CreateUseCase {
	return &createExpenseUseCase{
		srv: srv,
	}
}

func (u *createExpenseUseCase) CreateExpense(ctx context.Context, e *domain.Expense) (string, error) {
	res, err := u.srv.Create(ctx, e)
	if err != nil {
		return "", err
	}
	return res, nil
}
