package server

import (
	"fmt"
	"github.com/palindrom615/docker-exposer/logger"
	"net/http"
)

var log = logger.DefaultLogger()

type Server struct {
	server *http.Server
}

func NewServer(port int, handler http.Handler) *Server {
	return &Server{
		server: &http.Server{Addr: fmt.Sprintf(":%d", port), Handler: handler},
	}
}

func (s *Server) Start() {
	log.Infof("Server listening on %s", s.server.Addr)
	s.server.ListenAndServe()
}
