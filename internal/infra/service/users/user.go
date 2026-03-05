package users

import (
	"context"
	"profitti/internal/core/domain"
	"profitti/internal/infra/database/repository/user"

	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	Register(context.Context, *domain.User) (string, error)
	FindOne(context.Context, string) (*domain.User, error)
}

type service struct {
	repo user.UserRepo
}

func New(repo user.UserRepo) UserService {
	return &service{
		repo: repo,
	}
}

func HashPassword(password *string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(*password), 14)
	*password = string(bytes)
	return err
}

func (s *service) Register(ctx context.Context, user *domain.User) (string, error) {
	err := HashPassword(&user.Password)
	if err != nil {
		return "", err
	}
	res, err := s.repo.InsertOne(ctx, user)
	if err != nil {
		return "", err
	}
	return res, nil
}

func (s *service) FindOne(ctx context.Context, id string) (*domain.User, error) {
	res, err := s.repo.SelectOne(ctx, id)
	if err != nil {
		return nil, err
	}
	return res, nil
}
