# docker-exposer

## Description

This is a simple server that exposes the docker socket to the outside world through http.

Although [docker daemon already supports remote access via http](https://docs.docker.com/config/daemon/remote-access/),
built-in authentication method is very limited.

You can expose your docker socket with additional features built-in or with your own.

## Quick Start

**WARNING**

This project itself does not provide any authentication mechanism by default. **NEVER** expose this server to the
public internet without additional authentication mechanism.

### native

```bash
go install github.com/palindrom615/docker-exposer/cmd/docker-exposer@latest
~/go/bin/docker-exposer --port 8080 --docker-host unix://$HOME/.colima/docker.sock \
  --auth-type basic --basic-auth-username alice --basic-auth-password pa55word
```

### docker

```bash
DOCKER_HOST=/var/run/docker.sock
docker run -d -p 8080:8080 \
  -v $DOCKER_HOST:/var/run/docker.sock \
  -e AUTH_TYPE=basic \
  -e BASIC_AUTH_USERNAME=alice \
  -e BASIC_AUTH_PASSWORD=pa55word \
  --name docker-exposer \
  ghcr.io/palindrom615/docker-exposer:latest
```

## Options

.env file is supported.

| flag                  | description                                                                      | default                          |
|-----------------------|----------------------------------------------------------------------------------|----------------------------------|
| --port                | port number to listen                                                            | 8080                             |
| --docker-host         | docker host to connect                                                           | os.Getenv("DOCKER_HOST")         |
| --docker-ca-cert      | path to docker CA cert on tls verify                                             | nil                              |
| --docker-cert         | path to docker cert on tls verify                                                | nil                              |
| --docker-key          | path to docker key on tls verify                                                 | nil                              |
| --auth-type           | authentication type. `basic` and `auth0` is supported. anything else is ignored. | os.Getenv("AUTH_TYPE")           |
| --basic-auth-username | basic auth username                                                              | os.Getenv("BASIC_AUTH_USERNAME") |
| --basic-auth-password | basic auth password                                                              | os.Getenv("BASIC_AUTH_PASSWORD") |

### built-in authentication

#### basic auth

[Built-in basic auth implementation](middleware/basicauth.go) only supports single username and password given
by `--basic-auth-username` and `--basic-auth-password`.

#### auth0

[built-in auth0 implementation](middleware/auth0.go)
supports [auth0 Backend/API](https://auth0.com/docs/quickstart/backend/golang/01-authorization#configure-auth0-apis)
authentication.
this implementation requires environment variable `AUTH0_DOMAIN` for your auth0 tenant and `AUTH0_AUDIENCE` for your
auth0 api ID. 
