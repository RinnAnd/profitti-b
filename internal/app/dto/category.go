package dto

import "profitti/internal/core/domain"

type CategoryRes struct {
	UserId     string     `json:"user_id"`
	Categories []Category `json:"categories"`
}

type Category struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

func (c *Category) Domain() *domain.Category {
	return &domain.Category{
		Name: c.Name,
	}
}
