package main

import (
	"bufio"
	"context"
	"github.com/docker/docker/client"
	"io"
	"log"
	"log/slog"
	"math/rand"
	"net/http"
	"os"
)

var logLevel = new(slog.LevelVar)
var logger = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: logLevel}))
var RequestIDKey = "request_id"

func main() {
	logLevel.Set(slog.LevelDebug)
	logger.Enabled(context.Background(), slog.LevelDebug)
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		logger.Error("Failed to create Docker client", "err", err)
		os.Exit(1)
	}
	conn, err := cli.Dialer()(context.Background())
	if err != nil {
		logger.Error("Failed to get Docker connection:", err)
		os.Exit(1)
	}
	defer conn.Close()

	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		rid := rand.Uint64()
		logger := logger.With(RequestIDKey, rid)
		logger.Debug("Request", "req", req)
		if err := req.Write(conn); err != nil {
			logger.Error("Failed to write request:", err)
			return
		}
		res, err := http.ReadResponse(bufio.NewReader(conn), req)
		if err != nil {
			logger.Error("Failed to read response:", err)
			return
		}
		defer res.Body.Close()
		w.WriteHeader(res.StatusCode)
		for key, values := range res.Header {
			for _, value := range values {
				w.Header().Add(key, value)
			}
		}
		logger.Debug("Response", "res", res)
		body, err := io.ReadAll(res.Body)
		if err != nil {
			logger.Error("Failed to read response body:", err)
			return
		}
		w.Write(body)
	})

	log.Println("Server listening on :8080")
	http.ListenAndServe(":8080", nil)
}
