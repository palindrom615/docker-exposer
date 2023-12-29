package main

import (
	"context"
	"docker-exposer"
	"github.com/docker/docker/client"
	"io"
	"net/http"
	"os"
)

func main() {
	log := docker_exposer.DefaultLogger()
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		log.Error("Failed to create Docker client", "err", err)
		os.Exit(1)
	}
	conn, err := cli.Dialer()(context.Background())
	if err != nil {
		log.Error("Failed to get Docker connection:", err)
		os.Exit(1)
	}

	relay := docker_exposer.NewRelay(conn)
	roundTripper := docker_exposer.NewRequestLog(log, relay)
	defer conn.Close()

	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		res, err := roundTripper.RoundTrip(req)
		if err != nil {
			log.Error("Failed to read response:", err)
			return
		}
		defer res.Body.Close()
		w.WriteHeader(res.StatusCode)
		for key, values := range res.Header {
			for _, value := range values {
				w.Header().Add(key, value)
			}
		}
		body, err := io.ReadAll(res.Body)
		if err != nil {
			log.Error("Failed to read response body:", err)
			return
		}
		w.Write(body)
	})

	log.Info("Server listening on :8080")
	http.ListenAndServe(":8080", nil)
}
