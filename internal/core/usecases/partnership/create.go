package partnership

import (
	"context"
	"profitti/internal/core/domain"
	"profitti/internal/infra/service/partnership"
)

type usecases struct {
	srv partnership.Service
}

type UseCase interface {
	Create(context.Context, *domain.Partnership) (string, error)
	GetByUser(context.Context, string) ([]*domain.Partnership, error)
}

func New(srv partnership.Service) UseCase {
	return &usecases{
		srv: srv,
	}
}

func (u *usecases) Create(ctx context.Context, p *domain.Partnership) (string, error) {
	res, err := u.srv.Create(ctx, p)
	if err != nil {
		return "", err
	}

	return res, nil
}

func (u *usecases) GetByUser(ctx context.Context, id string) ([]*domain.Partnership, error) {
	res, err := u.srv.GetPartnerships(ctx, id)
	if err != nil {
		return nil, err
	}

	return res, nil
}
