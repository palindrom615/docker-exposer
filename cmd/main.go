package main

import (
	"context"
	"docker-exposer"
	"github.com/docker/docker/client"
	"io"
	"math/rand"
	"net/http"
	"os"
)

var RequestIDKey = "request_id"

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
	defer conn.Close()

	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		rid := rand.Uint64()
		log := log.With(RequestIDKey, rid)
		log.Debug("Request", "req", req)

		res, err := relay.RoundTrip(req)
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
		log.Debug("Response", "res", res)
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
