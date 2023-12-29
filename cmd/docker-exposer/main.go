package main

import (
	"flag"
	"github.com/palindrom615/docker-exposer/pkg/client"
	"github.com/palindrom615/docker-exposer/pkg/serve"
)

func main() {
	port := flag.Int("port", 8080, "port to listen on")
	dockerHost := flag.String("docker-host", "", "docker host")
	dockerCaCert := flag.String("docker-cert-path", "", "docker cert path")
	dockerCert := flag.String("docker-cert", "", "docker cert")
	dockerKey := flag.String("docker-key", "", "docker key")
	flag.Parse()

	dockerOption := client.NewDockerOptions(*dockerHost, *dockerCaCert, *dockerCert, *dockerKey)

	server := serve.NewServer(*port, dockerOption)
	server.Start()
}
