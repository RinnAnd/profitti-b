package partnership

import (
	"context"
	"database/sql"
	"encoding/json"
	"profitti/internal/core/domain"
	"sort"
)

type Partnership interface {
	Insert(context.Context, *domain.Partnership) (string, error)
	Update(context.Context, *domain.Partnership) (string, error)
	Select(context.Context, string) ([]*domain.Partnership, error)
	SelectOne(context.Context, string) (*domain.Partnership, error)
}

type repository struct {
	db *sql.DB
}

func New(db *sql.DB) Partnership {
	return &repository{
		db: db,
	}
}

func (r *repository) Insert(ctx context.Context, p *domain.Partnership) (string, error) {
	var target string

	sort.Strings(p.Users)

	users := map[string][]string{
		"users": p.Users,
	}
	usersJSON, err := json.Marshal(users)
	if err != nil {
		return "", err
	}

	err = r.db.QueryRowContext(ctx, `
		INSERT INTO
		  partnership (users, currency_id)
		VALUES
		  ($1, $2)
		RETURNING
  		  id
		`, usersJSON, p.CurrencyId).Scan(&target)
	if err != nil {
		return "", err
	}

	return target, nil
}

func (r *repository) Select(ctx context.Context, id string) ([]*domain.Partnership, error) {
	var partnerships []*domain.Partnership

	rows, err := r.db.QueryContext(ctx, `
		SELECT
		  id, users, currency_id
		FROM
		  partnership
		WHERE
		  users @> jsonb_build_object('users', jsonb_build_array($1::text));
		`, id)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		usersBytes := []byte{}
		p := domain.Partnership{}
		if err := rows.Scan(&p.Id, &usersBytes, &p.CurrencyId); err != nil {
			return nil, err
		}

		uMap := map[string][]string{}
		if err := json.Unmarshal(usersBytes, &uMap); err != nil {
			return nil, err
		}

		p.Users = uMap["users"]
		partnerships = append(partnerships, &p)
	}

	return partnerships, nil
}

func (r *repository) SelectOne(ctx context.Context, id string) (*domain.Partnership, error) {
	var target domain.Partnership
	ubytes := []byte{}

	err := r.db.QueryRowContext(ctx, `
		SELECT
		  id, users, currency_id
		FROM
		  partnership
		WHERE
		  id = $1;
		`, id).Scan(&target.Id, &ubytes, &target.CurrencyId)
	if err != nil {
		return nil, err
	}

	umap := map[string][]string{}

	if err = json.Unmarshal(ubytes, &umap); err != nil {
		return nil, err
	}

	target.Users = umap["users"]

	return &target, nil
}

func (r *repository) Update(ctx context.Context, p *domain.Partnership) (string, error) {
	var target string

	users := map[string][]string{
		"users": p.Users,
	}
	usersJSON, err := json.Marshal(users)
	if err != nil {
		return "", err
	}

	err = r.db.QueryRowContext(ctx, `
		UPDATE partnership
		SET
		  users = $1
		WHERE
		  id = $2
		RETURNING
		  id;
		`, usersJSON, p.Id).Scan(&target)
	if err != nil {
		return "", err
	}

	return target, nil
}
