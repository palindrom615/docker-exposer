package docker_exposer

import (
	"context"
	"github.com/docker/docker/client"
	"net"
)

var cli *client.Client

func init() {
	var err error
	cli, err = client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		panic(err)
	}
}

func GetDockerClient() *client.Client {
	if cli == nil {
		var err error
		cli, err = client.NewClientWithOpts(client.FromEnv)
		if err != nil {
			panic(err)
		}
	}
	return cli
}

func GetDockerConnection(ctx context.Context) net.Conn {
	conn, err := GetDockerClient().Dialer()(ctx)
	if err != nil {
		panic(err)
	}
	return conn
}
