package routes

import (
	"profitti/internal/app/transport/http/handlers/users"

	"github.com/gin-gonic/gin"
)

type handlers struct {
	userHandler users.UserHandler
}

func (h *handlers) Init(s *gin.Engine) {
	// Users
	userGroup := s.Group("/users")
	userGroup.POST("/register", h.userHandler.Register)
}
