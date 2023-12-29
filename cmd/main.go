package main

import (
	"context"
	"docker-exposer"
	"io"
	"net/http"
)

func main() {
	log := docker_exposer.DefaultLogger()
	conn := docker_exposer.GetDockerConnection(context.Background())

	relay := docker_exposer.NewRelay(conn)
	roundTripper := docker_exposer.NewRequestLog(log, relay)

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
