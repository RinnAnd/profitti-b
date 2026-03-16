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
	// var sharedExpenses *[]domain.Expense
	err := r.db.QueryRowContext(ctx,
		`INSERT INTO expenses (financial_id, partnership_id, name, description, amount, expense_recurrence, expiration_date, currency_id) VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id`,
		expense.Financial_id, expense.Partnership_id, expense.Name, expense.Description, expense.Amount, expense.Expense_recurrence, expense.Expiration_date, expense.Currency_id).Scan(&target)
	if err != nil {
		return "", err
	}
	if expense.Partnership_id != nil {

	}
	return target, nil
}

func (r *repo) SelectUserExpenses(ctx context.Context, id string) ([]*domain.Expense, error) {
	var target []*domain.Expense
	rows, err := r.db.QueryContext(ctx,
		`SELECT e.*,
		c1.name AS currency_name,
		se.percentage AS shared_percentage,
		c2.name AS shared_currency_name
		FROM
		expenses e
		INNER JOIN currencies c1 ON e.currency_id = c1.id
		LEFT JOIN financials f ON e.financial_id = f.id
		LEFT JOIN partnership p ON e.partnership_id = p.id
		LEFT JOIN shared_expenses se ON e.id = se.expense_id
		AND se.user_id = $1
		LEFT JOIN currencies c2 ON se.currency_id = c2.id
		WHERE
		f.user_id = $1
		OR p.users @> CAST($1 AS TEXT)::jsonb
		OR se.user_id = $1`, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		e := &domain.Expense{}
		rows.Scan(e.Financial_id, e.Partnership_id, e.Name, e.Description, e.Amount, e.Expense_recurrence, e.Expiration_date, e.Currency_id, e.Currency, e.SharedPercentage, e.SharedCurrency)
		target = append(target, e)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return target, nil
}
