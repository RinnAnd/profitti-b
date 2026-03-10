package users

import (
	"database/sql"
	"errors"
	"net/http"
	"profitti/internal/app/dto"
	"profitti/internal/core/domain"
	"profitti/internal/infra/service/users"

	"github.com/gin-gonic/gin"
)

type lgnhandler struct {
	srv users.UserService
}

type LoginHandler interface {
	Login(c *gin.Context)
}

func NewLogin(srv users.UserService) LoginHandler {
	return &lgnhandler{
		srv: srv,
	}
}

func (l *lgnhandler) Login(c *gin.Context) {
	req, err := decodeLoginRq(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.HttpError{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	res, err := l.srv.FindOne(c, req.Email, req.Password)
	if err != nil {
		e := new(domain.CredentialsError)
		if errors.As(err, &e) || errors.Is(err, sql.ErrNoRows) {
			c.JSON(http.StatusUnauthorized, dto.HttpError{
				Status:  http.StatusUnauthorized,
				Message: "wrong credentials",
			})
			return
		}
		c.JSON(http.StatusInternalServerError, dto.HttpError{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusAccepted, res)
}

func decodeLoginRq(c *gin.Context) (*dto.Login, error) {
	req := &dto.Login{}

	if err := c.ShouldBindJSON(req); err != nil {
		return nil, err
	}

	return req, nil
}
