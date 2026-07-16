package categories

import (
	"context"
	"database/sql"
	"fmt"
	"profitti/internal/core/domain"
)

type Categories interface {
	Insert(context.Context, *domain.Category) error
	SelectAll(context.Context, string) ([]*domain.Category, error)
}

type repository struct {
	db *sql.DB
}

func New(db *sql.DB) Categories {
	return &repository{
		db: db,
	}
}

func (r *repository) Insert(ctx context.Context, c *domain.Category) error {
	_, err := r.db.ExecContext(ctx, `
		INSERT INTO categories (name, user_id) VALUES ($1, $2)
	`, c.Name, c.UserId)
	if err != nil {
		return err
	}
	return nil
}

func (r *repository) SelectAll(ctx context.Context, user_id string) ([]*domain.Category, error) {
	var target []*domain.Category

	rows, err := r.db.QueryContext(ctx, `
		SELECT * FROM categories WHERE user_id = $1;
	`, user_id)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		cat := &domain.Category{}
		err := rows.Scan(&cat.Id, &cat.Name, &cat.UserId)
		if err != nil {
			return nil, fmt.Errorf("error scanning values in category: %w", err)
		}

		target = append(target, cat)
	}

	return target, nil
}
