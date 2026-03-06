package server

import "github.com/gin-gonic/gin"

type Server struct {
	addr string
	g    *gin.Engine
}

func StartServer(addr string) *Server {
	router := gin.Default()
	return &Server{
		addr: addr,
		g:    router,
	}
}

func (s *Server) Run() {
	err := s.g.Run(s.addr)
	if err != nil {
		panic(err.Error())
	}
}
