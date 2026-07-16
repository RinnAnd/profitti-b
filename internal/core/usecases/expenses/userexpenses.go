package expenses

import (
	"context"
	"profitti/internal/core/domain"
	"profitti/internal/infra/service/expenses"
	"profitti/internal/infra/service/users"
)

type GetByUserUseCase interface {
	GetExpensesByUser(context.Context, string) ([]*domain.Expense, error)
}

type getByUserUseCase struct {
	srv    expenses.ExpenseService
	usrsrv users.UserService
}

func NewGetByUserUseCase(srv expenses.ExpenseService, usrsrv users.UserService) GetByUserUseCase {
	return &getByUserUseCase{
		srv:    srv,
		usrsrv: usrsrv,
	}
}

func (u *getByUserUseCase) GetExpensesByUser(ctx context.Context, id string) ([]*domain.Expense, error) {
	if !u.usrsrv.CheckOne(ctx, id) {
		return nil, domain.User404
	}

	res, err := u.srv.GetUserExpenses(ctx, id)
	if err != nil {
		return nil, err
	}
	return res, nil
}
