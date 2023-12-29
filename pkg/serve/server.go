package serve

import (
	"context"
	"fmt"
	"github.com/palindrom615/docker-exposer/pkg/client"
	"net/http"
)

type Server struct {
	server *http.Server
}

func NewServer(port int, options *client.DockerOptions) *Server {
	dockerClient := client.NewDockerClientFactory(options).NewDockerClient()
	conn, err := dockerClient.Dialer()(context.Background())
	if err != nil {
		panic(err)
	}

	dockerConnector := client.NewConnector(conn)
	roundTripper := NewRequestLog(dockerConnector)
	handler := NewRoundTripHandler(roundTripper)
	return &Server{
		server: &http.Server{Addr: fmt.Sprintf(":%d", port), Handler: handler},
	}
}

func (s *Server) Start() {
	log.Infof("Server listening on %s", s.server.Addr)
	s.server.ListenAndServe()
}
