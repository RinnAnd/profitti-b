package auth

import (
	"errors"
	"profitti/internal/core/domain"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type auth struct {
	secret []byte
	ttl    time.Duration
}

type JWT interface {
	GenAccessToken(*domain.User) (string, error)
	Validate(string) (*jwt.MapClaims, error)
}

func New(secret string, ttl time.Duration) JWT {
	return &auth{
		secret: []byte(secret),
		ttl:    ttl,
	}
}

func (a *auth) GenAccessToken(user *domain.User) (string, error) {
	expiration := time.Now().Add(a.ttl)

	claims := jwt.MapClaims{
		"username": user.Username,
		"email":    user.Email,
		"exp":      expiration.Unix(),
		"iat":      time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signed, err := token.SignedString(a.secret)
	if err != nil {
		return "", err
	}

	return signed, nil
}

func (a *auth) Validate(tk string) (*jwt.MapClaims, error) {
	token, err := jwt.Parse(tk, func(t *jwt.Token) (any, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid token")
		}
		return a.secret, nil
	})

	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			return nil, errors.New("expired token")
		}
		return nil, errors.New("invalid token")
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return &claims, nil
	}

	return nil, errors.New("invalid token")
}
