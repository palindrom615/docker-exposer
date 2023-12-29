package main

import (
	"github.com/palindrom615/docker-exposer/pkg/serve"
)

func main() {
	server := serve.NewServer(8080)
	server.Start()
}
