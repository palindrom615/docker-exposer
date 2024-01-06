package server

import (
	"flag"
	"github.com/joho/godotenv"
	"os"
)

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
	// AuthType is the auth type. `basic`, `auth0` is supported. anything else is ignored.
	// default: os.Getenv("AUTH_TYPE")
	AuthType string
	// BasicAuthUsername is the username for basic auth
	BasicAuthUsername string
	// BasicAuthPassword is the password for basic auth
	BasicAuthPassword string
}

func newFlags() *flags {
	_ = godotenv.Load()
	f := &flags{}

	flag.IntVar(&f.Port, "port", 2375, "port to listen on")
	flag.StringVar(&f.DockerHost, "docker-host", "", "docker host")
	flag.StringVar(&f.DockerCaCert, "docker-tlscacert", "", "docker cert path")
	flag.StringVar(&f.DockerCert, "docker-tlscert", "", "docker cert")
	flag.StringVar(&f.DockerKey, "docker-tlskey", "", "docker key")
	flag.StringVar(&f.AuthType, "auth", os.Getenv("AUTH_TYPE"), "auth type")
	flag.StringVar(&f.BasicAuthUsername, "basic-auth-username", os.Getenv("BASIC_AUTH_USERNAME"), "basic auth username")
	flag.StringVar(&f.BasicAuthPassword, "basic-auth-password", os.Getenv("BASIC_AUTH_PASSWORD"), "basic auth password")

	flag.Parse()
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

type authFlags interface {
	GetAuthType() string
	GetBasicAuthUsername() string
	GetBasicAuthPassword() string
}

func (f *flags) GetAuthType() string {
	return f.AuthType
}

func (f *flags) GetBasicAuthUsername() string {
	return f.BasicAuthUsername
}

func (f *flags) GetBasicAuthPassword() string {
	return f.BasicAuthPassword
}
