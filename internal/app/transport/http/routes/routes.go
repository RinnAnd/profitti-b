package routes

import (
	"profitti/internal/app/transport/http/handlers/users"

	"github.com/gin-gonic/gin"
)

type Routes struct {
	RegisterHandler users.RegisterHandler
	LoginHandler    users.LoginHandler
}

func (h *Routes) Init(s *gin.Engine) {
	userGroup := s.Group("/users")
	userGroup.POST("/register", h.RegisterHandler.Register)
	userGroup.POST("/login", h.LoginHandler.Login)
}
