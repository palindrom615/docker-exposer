package connector

import (
	"github.com/docker/docker/client"
)

func NewDockerOptions(host, cacert, cert, key string) *[]client.Opt {
	o := []client.Opt{client.FromEnv}
	if host != "" {
		o = append(o, client.WithHost(host))
	}
	if cacert != "" && cert != "" && key != "" {
		o = append(o, client.WithTLSClientConfig(cacert, cert, key))
	}
	return &o
}

func NewDockerClient(opts *[]client.Opt) client.CommonAPIClient {
	cli, err := client.NewClientWithOpts(*opts...)
	if err != nil {
		panic(err)
	}
	return cli
}
