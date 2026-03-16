package partnership

import (
	"context"
	"errors"
	"profitti/internal/core/domain"
	"profitti/internal/infra/database/repository/partnership"
	"slices"
)

type Service interface {
	Create(context.Context, *domain.Partnership) (string, error)
	GetPartnership(context.Context, string) (*domain.Partnership, error)
	GetPartnerships(context.Context, string) ([]*domain.Partnership, error)
	Update(context.Context, *domain.Partnership) (string, error)
}

type service struct {
	repo partnership.Partnership
}

func New(repo partnership.Partnership) Service {
	return &service{
		repo: repo,
	}
}

func (s *service) Create(ctx context.Context, p *domain.Partnership) (string, error) {
	res, err := s.repo.Insert(ctx, p)
	if err != nil {
		return "", err
	}
	return res, nil
}

func (s *service) GetPartnership(ctx context.Context, id string) (*domain.Partnership, error) {
	res, err := s.repo.SelectOne(ctx, id)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *service) GetPartnerships(ctx context.Context, id string) ([]*domain.Partnership, error) {
	res, err := s.repo.Select(ctx, id)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *service) Update(ctx context.Context, p *domain.Partnership) (string, error) {
	prev, err := s.GetPartnership(ctx, p.Id)
	if err != nil {
		return "", err
	}

	curr, err := validateFields(prev, p)
	if err != nil {
		return "", err
	}

	res, err := s.Update(ctx, curr)
	if err != nil {
		return "", err
	}
	return res, nil
}

func validateFields(prev, curr *domain.Partnership) (*domain.Partnership, error) {
	if !slices.Equal(prev.Users, curr.Users) {
		return curr, nil
	}
	return nil, errors.New("no change in users")
}
