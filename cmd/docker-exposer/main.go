package main

import (
	"docker-exposer/pkg/client"
	"docker-exposer/pkg/logger"
	"docker-exposer/pkg/serve"
	"net/http"
)

var log = logger.DefaultLogger()

func main() {
	dockerClient := client.NewDockerClient()
	roundTripper := serve.NewRequestLog(dockerClient)

	handler := serve.NewRoundTripHandler(roundTripper)
	http.Handle("/", handler)

	log.Info("Server listening on :8080")
	http.ListenAndServe(":8080", nil)
}
