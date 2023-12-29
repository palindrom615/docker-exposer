package main

import (
	"context"
	"docker-exposer"
	"net/http"
)

func main() {
	log := docker_exposer.DefaultLogger()
	conn := docker_exposer.GetDockerConnection(context.Background())

	relay := docker_exposer.NewRelay(conn)
	roundTripper := docker_exposer.NewRequestLog(log, relay)

	handler := docker_exposer.NewRoundTripHandler(roundTripper)
	http.Handle("/", handler)

	log.Info("Server listening on :8080")
	http.ListenAndServe(":8080", nil)
}
