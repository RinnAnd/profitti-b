package server

import (
	"profitti/internal/app/middleware"
	"profitti/internal/infra/service/auth"

	"github.com/gin-gonic/gin"
)

type Server struct {
	addr string
	G    *gin.Engine
	jwt  auth.JWT
}

func StartServer(addr string, jwt auth.JWT) *Server {
	router := gin.Default()

	authmidd := middleware.New(jwt)
	router.Use(authmidd.CheckToken())

	return &Server{
		addr: addr,
		G:    router,
	}
}

func (s *Server) Run() {
	err := s.G.Run(s.addr)
	if err != nil {
		panic(err.Error())
	}
}
