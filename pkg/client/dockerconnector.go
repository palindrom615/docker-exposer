package client

import (
	"github.com/docker/docker/client"
)

type DockerOptions struct {
	host   string
	cacert string
	cert   string
	key    string
}

type DockerClientFactory struct {
	opts *DockerOptions
}

func NewDockerClientFactory(opts *DockerOptions) *DockerClientFactory {
	return &DockerClientFactory{opts: opts}
}

func NewDockerOptions(host, cacert, cert, key string) *DockerOptions {
	return &DockerOptions{
		host:   host,
		cacert: cacert,
		cert:   cert,
		key:    key,
	}
}

func (d *DockerClientFactory) getOpts() []client.Opt {
	o := []client.Opt{client.FromEnv}
	if d.opts == nil {
		return o
	}
	if d.opts.host != "" {
		o = append(o, client.WithHost(d.opts.host))
	}
	if d.opts.cacert != "" && d.opts.cert != "" && d.opts.key != "" {
		o = append(o, client.WithTLSClientConfig(d.opts.cacert, d.opts.cert, d.opts.key))
	}
	return o
}

func (d *DockerClientFactory) NewDockerClient() client.CommonAPIClient {
	cli, err := client.NewClientWithOpts(d.getOpts()...)
	if err != nil {
		panic(err)
	}
	return cli
}
