package login

import (
	"context"
	"profitti/internal/core/domain"
	"profitti/internal/infra/service/auth"
	"profitti/internal/infra/service/users"
)

type LoginUseCase interface {
	Login(context.Context, string, string) (*domain.User, error)
	GenToken(*domain.User) (string, error)
}

type useCase struct {
	usersrv users.UserService
	auth    auth.JWT
}

func New(usersrv users.UserService, auth auth.JWT) LoginUseCase {
	return &useCase{
		usersrv: usersrv,
		auth:    auth,
	}
}

func (u *useCase) Login(ctx context.Context, email, password string) (*domain.User, error) {
	res, err := u.usersrv.FindOne(ctx, email, password)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (u *useCase) GenToken(user *domain.User) (string, error) {
	token, err := u.auth.GenAccessToken(user)
	if err != nil {
		return "", err
	}

	return token, nil
}
