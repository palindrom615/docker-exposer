package main

import (
	"github.com/palindrom615/docker-exposer/pkg/client"
	"github.com/palindrom615/docker-exposer/pkg/logger"
	"github.com/palindrom615/docker-exposer/pkg/serve"
	"net/http"
)

var log = logger.DefaultLogger()

func main() {
	dockerConnector := client.NewDockerConnector()
	roundTripper := serve.NewRequestLog(dockerConnector)

	handler := serve.NewRoundTripHandler(roundTripper)
	http.Handle("/", handler)

	log.Info("Server listening on :8080")
	http.ListenAndServe(":8080", nil)
}
