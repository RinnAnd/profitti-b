package categories

import (
	"context"
	"profitti/internal/core/domain"
	"profitti/internal/infra/service/categories"
	"profitti/internal/infra/service/users"
)

type createusecase struct {
	usr users.UserService
	cat categories.CatService
}

type CUseCase interface {
	CreateCategory(context.Context, *domain.Category) error
}

func New(usr users.UserService, cat categories.CatService) CUseCase {
	return &createusecase{
		usr: usr,
		cat: cat,
	}
}

func (u *createusecase) CreateCategory(ctx context.Context, c *domain.Category) error {
	if !u.usr.CheckOne(ctx, c.UserId) {
		return domain.User404
	}

	return u.cat.CreateCategory(ctx, c)
}
