package main

import (
	"github.com/palindrom615/docker-exposer/server"
)

func main() {
	s := server.Configure()
	s.Start()
}
