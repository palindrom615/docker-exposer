package client

import (
	"bufio"
	"github.com/docker/docker/client"
	"net/http"
)

type DockerClient struct {
	cli client.CommonAPIClient
}

func NewDockerClient() *DockerClient {
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		panic(err)
	}
	return &DockerClient{cli: cli}
}

func (d *DockerClient) RoundTrip(req *http.Request) (*http.Response, error) {
	conn, err := d.cli.Dialer()(req.Context())
	if err != nil {
		return nil, err
	}
	if err := req.Write(conn); err != nil {
		return nil, err
	}
	return http.ReadResponse(bufio.NewReader(conn), req)
}
