package main

import (
	"bufio"
	"context"
	"github.com/docker/docker/client"
	"io"
	"log"
	"net/http"
)

var logger = log.Default()

func main() {
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		logger.Panicln("Failed to create Docker client:", err)
	}
	conn, err := cli.Dialer()(context.Background())
	if err != nil {
		logger.Panicln("Failed to get Docker connection:", err)
	}
	defer conn.Close()

	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		req.Write(logger.Writer())
		if err := req.Write(conn); err != nil {
			logger.Fatalln("Failed to write request:", err)
			return
		}
		res, err := http.ReadResponse(bufio.NewReader(conn), req)
		if err != nil {
			logger.Fatalln("Failed to read response:", err)
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
			logger.Fatalln("Failed to read response body:", err)
			return
		}
		w.Header().Set("Content-Length", string(len(body)))
		w.Write(body)
	})

	log.Println("Server listening on :8080")
	http.ListenAndServe(":8080", nil)
}
