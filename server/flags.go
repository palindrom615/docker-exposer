package server

import "flag"

type flags struct {
	// Port is the port to listen on
	Port int
	// DockerHost is the docker host to connect to
	DockerHost string
	// DockerCaCert is the path to the docker ca cert
	DockerCaCert string
	// DockerCert is the path to the docker cert
	DockerCert string
	// DockerKey is the path to the docker key
	DockerKey string
}

func newFlags() *flags {
	var port = flag.Int("port", 8080, "port to listen on")
	var dockerHost = flag.String("docker-host", "", "docker host")
	var dockerCaCert = flag.String("docker-cert-path", "", "docker cert path")
	var dockerCert = flag.String("docker-cert", "", "docker cert")
	var dockerKey = flag.String("docker-key", "", "docker key")
	flag.Parse()
	f := &flags{
		Port:         *port,
		DockerHost:   *dockerHost,
		DockerCaCert: *dockerCaCert,
		DockerCert:   *dockerCert,
		DockerKey:    *dockerKey,
	}
	return f
}

type serverFlags interface {
	GetPort() int
}

func (f *flags) GetPort() int {
	return f.Port
}

type dockerFlags interface {
	GetDockerHost() string
	GetDockerCaCert() string
	GetDockerCert() string
	GetDockerKey() string
}

func (f *flags) GetDockerHost() string {
	return f.DockerHost
}

func (f *flags) GetDockerCaCert() string {
	return f.DockerCaCert
}

func (f *flags) GetDockerCert() string {
	return f.DockerCert
}

func (f *flags) GetDockerKey() string {
	return f.DockerKey
}
