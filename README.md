# docker-exposer

## Description

This is a simple server that exposes the docker socket to the outside world through http.

Although [docker daemon already supports remote access via http](https://docs.docker.com/config/daemon/remote-access/),
built-in authentication method is limited to TLS.

You can modify this project and serve docker socket with additional features like authentication or CORS filter.

## Usage

**WARNING** 

This project itself does not provide any authentication mechanism. **NEVER** expose this server to the
public internet without additional authentication mechanism.

### native

```bash
go install github.com/palindrom615/docker-exposer/cmd/docker-exposer@latest
~/go/bin/docker-exposer --port 8080 --docker-host unix://$HOME/.colima/docker.sock
```

### docker

```bash
DOCKER_HOST=/var/run/docker.sock
docker run -d -p 8080:8080 -v $DOCKER_HOST:/var/run/docker.sock --name docker-exposer ghcr.io/palindrom615/docker-exposer:latest
```
