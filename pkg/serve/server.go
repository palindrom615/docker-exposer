package serve

import (
	"fmt"
	"github.com/palindrom615/docker-exposer/pkg/client"
	"net/http"
)

type Server struct {
	port   int
	server *http.Server
}

func NewServer(port int) *Server {
	dockerConnector := client.NewDockerConnector()
	roundTripper := NewRequestLog(dockerConnector)
	handler := NewRoundTripHandler(roundTripper)
	return &Server{
		port:   port,
		server: &http.Server{Addr: fmt.Sprintf(":%d", port), Handler: handler},
	}
}

func (s *Server) Start() {
	log.Info("Server listening on :8080")
	s.server.ListenAndServe()
}
