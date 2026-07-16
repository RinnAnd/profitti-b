package expense

import (
	"context"
	"database/sql"
	"profitti/internal/core/domain"
)

type ExpenseRepo interface {
	InsertOne(context.Context, *domain.Expense) (string, error)
	SelectUserExpenses(context.Context, string) ([]*domain.Expense, error)
}

type repo struct {
	db *sql.DB
}

func New(db *sql.DB) ExpenseRepo {
	return &repo{
		db: db,
	}
}

func (r *repo) InsertOne(ctx context.Context, expense *domain.Expense) (string, error) {
	var target string
	err := r.db.QueryRowContext(ctx, `
		INSERT INTO
			expenses (
				financial_id,
				partnership_id,
				name,
				description,
				amount,
				category_id,
				expense_recurrence,
				expiration_date,
				currency_id
			)
		VALUES
			($1, $2, $3, $4, $5, $6, $7, $8, $9)
		RETURNING
			id
	`, expense.FinancialId,
		expense.PartnershipId,
		expense.Name,
		expense.Description,
		expense.Amount,
		expense.Category,
		expense.Expense_recurrence,
		expense.Expiration_date,
		expense.CurrencyId).Scan(&target)
	if err != nil {
		return "", err
	}

	return target, nil
}

func (r *repo) SelectUserExpenses(ctx context.Context, id string) ([]*domain.Expense, error) {
	var target []*domain.Expense
	rows, err := r.db.QueryContext(ctx, `
		SELECT 
			e.id,
			e.financial_id,
			e.partnership_id,
			e.name,
			e.description,
			e.amount,
			c.name AS category,
			rt.type AS expense_recurrence,
			e.expiration_date,
			e.currency_id,
			e.created_at
		FROM expenses e
		LEFT JOIN recurrence_type rt ON e.expense_recurrence = rt.id
		LEFT JOIN categories c ON e.category_id  = c.id
		WHERE e.financial_id = $1 OR e.partnership_id = $1;
	`, id)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		expense := &domain.Expense{}

		err := rows.Scan(
			&expense.Id,
			&expense.FinancialId,
			&expense.PartnershipId,
			&expense.Name,
			&expense.Description,
			&expense.Amount,
			&expense.Category,
			&expense.Expense_recurrence,
			&expense.Expiration_date,
			&expense.CurrencyId,
			&expense.CreatedAt,
		)
		if err != nil {
			return nil, err
		}

		target = append(target, expense)
	}

	return target, nil
}
