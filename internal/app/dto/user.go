package dto

import "profitti/internal/core/domain"

type User struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Profile  string `json:"profile"`
}

func (u *User) Domain() *domain.User {
	return &domain.User{
		Username: u.Username,
		Email:    u.Email,
		Password: u.Password,
		Profile:  u.Profile,
	}
}

type UserRes struct {
	Id       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Profile  string `json:"profile"`
}
