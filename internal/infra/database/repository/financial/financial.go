package financial

import (
	"context"
	"database/sql"
	"profitti/internal/core/domain"
)

type FinancialRepo interface {
	InsertOne(context.Context, *domain.Financial) (string, error)
	SelectUserFinancials(context.Context, string) ([]*domain.Financial, error)
}

type repo struct {
	db *sql.DB
}

func New(db *sql.DB) FinancialRepo {
	return &repo{
		db: db,
	}
}

func (f *repo) InsertOne(ctx context.Context, financial *domain.Financial) (string, error) {
	var target string
	err := f.db.QueryRowContext(ctx, `
	INSERT INTO financials (user_id, currency_id) VALUES ($1, $2) RETURNING id`, financial.UserId, financial.CurrencyId).Scan(&target)
	if err != nil {
		return "", err
	}
	return target, nil
}

func (f *repo) SelectUserFinancials(ctx context.Context, id string) ([]*domain.Financial, error) {
	var target []*domain.Financial
	rows, err := f.db.QueryContext(ctx, `SELECT * FROM financials WHERE user_id = $1`, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		f := &domain.Financial{}
		rows.Scan(&f.Id, &f.UserId, &f.CurrencyId)
		target = append(target, f)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return target, nil
}
