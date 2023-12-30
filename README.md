# docker-exposer

## Description

This is a simple server that exposes the docker socket to the outside world through http.

Although [docker daemon already supports remote access via http](https://docs.docker.com/config/daemon/remote-access/),
built-in authentication method is limited to TLS.

You can modify this project and serve docker socket with additional features like authentication or CORS filter.

## Quick Start

**WARNING**

This project itself does not provide any authentication mechanism by default. **NEVER** expose this server to the
public internet without additional authentication mechanism.

### native

```bash
go install github.com/palindrom615/docker-exposer/cmd/docker-exposer@latest
~/go/bin/docker-exposer --port 8080 --docker-host unix://$HOME/.colima/docker.sock --enable-basic-auth --basic-auth-username alice --basic-auth-password pa55word
```

### docker

```bash
DOCKER_HOST=/var/run/docker.sock
docker run -d -p 8080:8080 -v $DOCKER_HOST:/var/run/docker.sock --name docker-exposer ghcr.io/palindrom615/docker-exposer:latest
```

## Options

| flag                  | description                          | default                          |
|-----------------------|--------------------------------------|----------------------------------|
| --port                | port number to listen                | 8080                             |
| --docker-host         | docker host to connect               | os.Getenv("DOCKER_HOST")         |
| --docker-ca-cert      | path to docker CA cert on tls verify | nil                              |
| --docker-cert         | path to docker cert on tls verify    | nil                              |
| --docker-key          | path to docker key on tls verify     | nil                              |
| --enable-basic-auth   | enable basic auth                    | false                            |
| --basic-auth-username | basic auth username                  | os.Getenv("BASIC_AUTH_USERNAME") |
| --basic-auth-password | basic auth password                  | os.Getenv("BASIC_AUTH_PASSWORD") |
