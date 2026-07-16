package partnership

import (
	"context"
	"profitti/internal/core/domain"
	"profitti/internal/infra/service/partnership"
	"profitti/internal/infra/service/users"
)

type usecases struct {
	srv    partnership.Service
	usrsrv users.UserService
}

type UseCase interface {
	Create(context.Context, *domain.Partnership) (string, error)
	GetByUser(context.Context, string) ([]*domain.Partnership, error)
}

func New(srv partnership.Service, usrsrv users.UserService) UseCase {
	return &usecases{
		srv:    srv,
		usrsrv: usrsrv,
	}
}

func (u *usecases) Create(ctx context.Context, p *domain.Partnership) (string, error) {
	for _, user := range p.Users {
		if !u.usrsrv.CheckOne(ctx, user) {
			return "", domain.User404
		}
	}

	return u.srv.Create(ctx, p)
}

func (u *usecases) GetByUser(ctx context.Context, id string) ([]*domain.Partnership, error) {
	if u.usrsrv.CheckOne(ctx, id) {
		return u.srv.GetPartnerships(ctx, id)
	}

	return nil, domain.User404
}
