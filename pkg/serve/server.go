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

func NewServer(port int, options *client.DockerOptions) *Server {
	dockerConnector := client.NewDockerConnector(options)
	roundTripper := NewRequestLog(dockerConnector)
	handler := NewRoundTripHandler(roundTripper)
	return &Server{
		port:   port,
		server: &http.Server{Addr: fmt.Sprintf(":%d", port), Handler: handler},
	}
}

func (s *Server) Start() {
	log.Infof("Server listening on :%d", s.port)
	s.server.ListenAndServe()
}
