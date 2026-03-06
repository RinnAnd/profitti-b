package server

import "github.com/gin-gonic/gin"

type Server struct {
	addr string
	G    *gin.Engine
}

func StartServer(addr string) *Server {
	router := gin.Default()
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
