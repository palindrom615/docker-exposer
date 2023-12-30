package server

import (
	"context"
	"github.com/docker/docker/client"
	"github.com/palindrom615/docker-exposer/middleware"
	"net/http"
)

func Configure() *Server {
	flags := provideFlags()
	port := providePort(flags)
	dockerOption := provideDockerOptions(flags)
	dockerClient := provideDockerClient(dockerOption)
	core := provideCore(dockerClient)
	middlewares := provideMiddlewares(flags)
	handler := provideHandler(core, middlewares)

	return NewServer(port, handler)
}

func provideFlags() *flags {
	return newFlags()
}

func providePort(flags serverFlags) int {
	return flags.GetPort()
}

func provideDockerOptions(flags dockerFlags) *[]client.Opt {
	dockerHost := flags.GetDockerHost()
	dockerCaCert := flags.GetDockerCaCert()
	dockerCert := flags.GetDockerCert()
	dockerKey := flags.GetDockerKey()

	o := []client.Opt{client.FromEnv}
	if dockerHost != "" {
		o = append(o, client.WithHost(dockerHost))
	}
	if dockerCaCert != "" && dockerCert != "" && dockerKey != "" {
		o = append(o, client.WithTLSClientConfig(dockerCaCert, dockerCert, dockerKey))
	}
	return &o
}

func provideDockerClient(opts *[]client.Opt) client.CommonAPIClient {
	cli, err := client.NewClientWithOpts(*opts...)
	if err != nil {
		panic(err)
	}
	return cli
}

func provideMiddlewares(flags basicAuthFlags) []middleware.Middleware {
	if flags.GetEnableBasicAuth() && flags.GetBasicAuthUsername() != "" && flags.GetBasicAuthPassword() != "" {
		return []middleware.Middleware{
			middleware.NewBasicAuth(flags.GetBasicAuthUsername(), flags.GetBasicAuthPassword()).Middleware,
			middleware.RequestLogHandler,
		}
	}
	return []middleware.Middleware{
		middleware.RequestLogHandler,
	}
}

func provideCore(dockerClient client.CommonAPIClient) http.Handler {
	conn, err := dockerClient.Dialer()(context.Background())
	if err != nil {
		panic(err)
	}

	return NewConnector(conn)
}

func provideHandler(core http.Handler, middlewares []middleware.Middleware) http.Handler {
	return middleware.ChainHandlers(core, middlewares...)
}