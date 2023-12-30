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
	// EnableBasicAuth enables basic auth
	EnableBasicAuth bool
	// BasicAuthUsername is the username for basic auth
	BasicAuthUsername string
	// BasicAuthPassword is the password for basic auth
	BasicAuthPassword string
}

func newFlags() *flags {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading the .env file: %v", err)
	}

	var port = flag.Int("port", 8080, "port to listen on")
	var dockerHost = flag.String("docker-host", "", "docker host")
	var dockerCaCert = flag.String("docker-ca-cert", "", "docker cert path")
	var dockerCert = flag.String("docker-cert", "", "docker cert")
	var dockerKey = flag.String("docker-key", "", "docker key")
	var enableBasicAuth = flag.Bool("enable-basic-auth", false, "enable basic auth")
	var basicAuthUsername = flag.String("basic-auth-username", os.Getenv("BASIC_AUTH_USERNAME"), "basic auth username")
	var basicAuthPassword = flag.String("basic-auth-password", os.Getenv("BASIC_AUTH_PASSWORD"), "basic auth password")

	flag.Parse()
	f := &flags{
		Port:              *port,
		DockerHost:        *dockerHost,
		DockerCaCert:      *dockerCaCert,
		DockerCert:        *dockerCert,
		DockerKey:         *dockerKey,
		EnableBasicAuth:   *enableBasicAuth,
		BasicAuthUsername: *basicAuthUsername,
		BasicAuthPassword: *basicAuthPassword,
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

type basicAuthFlags interface {
	GetEnableBasicAuth() bool
	GetBasicAuthUsername() string
	GetBasicAuthPassword() string
}

func (f *flags) GetEnableBasicAuth() bool {
	return f.EnableBasicAuth
}

func (f *flags) GetBasicAuthUsername() string {
	return f.BasicAuthUsername
}

func (f *flags) GetBasicAuthPassword() string {
	return f.BasicAuthPassword
}
