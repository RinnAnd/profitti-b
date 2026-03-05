package user

import (
	"context"
	"database/sql"
	"profitti/internal/core/domain"
)

type UserRepo interface {
	InsertOne(context.Context, *domain.User) (string, error)
	SelectOne(context.Context, string) (*domain.User, error)
}

type repo struct {
	db *sql.DB
}

func New(db *sql.DB) UserRepo {
	return &repo{
		db: db,
	}
}

func (u *repo) InsertOne(ctx context.Context, user *domain.User) (string, error) {
	var target string
	err := u.db.QueryRowContext(ctx, "INSERT INTO users (username, email, password, profile) VALUES ($1, $2, $3, $4) RETURNING id", user.Username, user.Email, user.Password, user.Profile).Scan(&target)
	if err != nil {
		return "", err
	}
	return target, nil
}

func (u *repo) SelectOne(ctx context.Context, id string) (*domain.User, error) {
	var target domain.User
	err := u.db.QueryRowContext(ctx, "SELECT id, username, email, profile FROM users WHERE id = $1", id).Scan(&target.Id, &target.Username, &target.Email, &target.Profile)
	if err != nil {
		return nil, err
	}
	return &target, nil
}
