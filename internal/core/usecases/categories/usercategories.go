package categories

import (
	"context"
	"profitti/internal/core/domain"
	"profitti/internal/infra/service/categories"
	"profitti/internal/infra/service/users"
)

type getusecase struct {
	usr users.UserService
	cat categories.CatService
}

type GUseCase interface {
	GetUserCategories(context.Context, string) ([]*domain.Category, error)
}

func NewG(usr users.UserService, cat categories.CatService) GUseCase {
	return &getusecase{
		usr: usr,
		cat: cat,
	}
}

func (g *getusecase) GetUserCategories(ctx context.Context, id string) ([]*domain.Category, error) {
	if !g.usr.CheckOne(ctx, id) {
		return nil, domain.User404
	}

	return g.cat.GetUserCategories(ctx, id)
}
