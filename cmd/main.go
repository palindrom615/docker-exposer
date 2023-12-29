package main

import (
	"context"
	"docker-exposer"
	"net/http"
)

var log = docker_exposer.DefaultLogger()

func main() {
	conn := docker_exposer.GetDockerConnection(context.Background())

	relay := docker_exposer.NewRelay(conn)
	roundTripper := docker_exposer.NewRequestLog(relay)

	handler := docker_exposer.NewRoundTripHandler(roundTripper)
	http.Handle("/", handler)

	log.Info("Server listening on :8080")
	http.ListenAndServe(":8080", nil)
}
