package server

import (
	"context"
	"fmt"
	"github.com/docker/docker/client"
	"github.com/palindrom615/docker-exposer/connector"
	"github.com/palindrom615/docker-exposer/logger"
	"github.com/palindrom615/docker-exposer/middleware"
	"net/http"
)

var log = logger.DefaultLogger()

type Server struct {
	server *http.Server
}

func NewServer(port int, options *[]client.Opt) *Server {
	dockerClient := connector.NewDockerClient(options)
	conn, err := dockerClient.Dialer()(context.Background())
	if err != nil {
		panic(err)
	}

	dockerConnector := connector.NewConnector(conn)
	handler := middleware.NewHandler(dockerConnector, middleware.RequestLog)
	return &Server{
		server: &http.Server{Addr: fmt.Sprintf(":%d", port), Handler: handler},
	}
}

func (s *Server) Start() {
	log.Infof("Server listening on %s", s.server.Addr)
	s.server.ListenAndServe()
}
