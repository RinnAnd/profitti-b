package routes

import (
	"profitti/internal/app/transport/http/handlers/users"

	"github.com/gin-gonic/gin"
)

type Routes struct {
	UserHandler users.UserHandler
}

func (h *Routes) Init(s *gin.Engine) {
	// Users
	userGroup := s.Group("/users")
	userGroup.POST("/register", h.UserHandler.Register)
}
