package client

import (
	"bufio"
	"context"
	"github.com/docker/docker/client"
	"net"
	"net/http"
)

type DockerConnector struct {
	conn net.Conn
}

func NewDockerConnector() *DockerConnector {
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		panic(err)
	}
	conn, err := cli.Dialer()(context.Background())
	if err != nil {
		panic(err)
	}
	return &DockerConnector{conn: conn}
}

func (d *DockerConnector) RoundTrip(req *http.Request) (*http.Response, error) {
	if err := req.Write(d.conn); err != nil {
		return nil, err
	}
	return http.ReadResponse(bufio.NewReader(d.conn), req)
}
