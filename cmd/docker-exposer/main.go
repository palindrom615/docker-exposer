package main

import (
	"docker-exposer"
	"docker-exposer/pkg/client"
	"docker-exposer/pkg/logger"
	"net/http"
)

var log = logger.DefaultLogger()

func main() {
	dockerClient := client.NewDockerClient()
	roundTripper := docker_exposer.NewRequestLog(dockerClient)

	handler := docker_exposer.NewRoundTripHandler(roundTripper)
	http.Handle("/", handler)

	log.Info("Server listening on :8080")
	http.ListenAndServe(":8080", nil)
}
