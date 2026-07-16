package categories

import (
	"context"
	"profitti/internal/core/domain"
	"profitti/internal/infra/database/repository/categories"
)

type service struct {
	repo categories.Categories
}

type CatService interface {
	GetUserCategories(context.Context, string) ([]*domain.Category, error)
	CreateCategory(context.Context, *domain.Category) error
}

func New(repo categories.Categories) CatService {
	return &service{
		repo: repo,
	}
}

func (s *service) GetUserCategories(ctx context.Context, id string) ([]*domain.Category, error) {
	return s.repo.SelectAll(ctx, id)
}

func (s *service) CreateCategory(ctx context.Context, c *domain.Category) error {
	return s.repo.Insert(ctx, c)
}
