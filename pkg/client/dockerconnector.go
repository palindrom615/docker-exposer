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

type DockerOptions struct {
	host   string
	cacert string
	cert   string
	key    string
}

func NewDockerOptions(host, cacert, cert, key string) *DockerOptions {
	return &DockerOptions{
		host:   host,
		cacert: cacert,
		cert:   cert,
		key:    key,
	}
}

func NewDockerConnector(opts *DockerOptions) *DockerConnector {
	o := []client.Opt{client.FromEnv}
	if opts.host != "" {
		o = append(o, client.WithHost(opts.host))
	}
	if opts.cacert != "" && opts.cert != "" && opts.key != "" {
		o = append(o, client.WithTLSClientConfig(opts.cacert, opts.cert, opts.key))
	}
	cli, err := client.NewClientWithOpts(o...)
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
