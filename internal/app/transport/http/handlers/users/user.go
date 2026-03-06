package users

import (
	"fmt"
	"net/http"
	"profitti/internal/app/dto"
	"profitti/internal/infra/service/users"

	"github.com/gin-gonic/gin"
)

type handler struct {
	srv users.UserService
}

type UserHandler interface {
	Register(c *gin.Context)
}

func NewUserHandler(srv users.UserService) UserHandler {
	return &handler{
		srv: srv,
	}
}

func (h *handler) Register(c *gin.Context) {
	req, err := decodeRequest(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.HttpError{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	fmt.Println(req)
}

func decodeRequest(c *gin.Context) (*dto.User, error) {
	req := &dto.User{}

	if err := c.ShouldBindJSON(req); err != nil {
		return nil, err
	}

	return req, nil
}
