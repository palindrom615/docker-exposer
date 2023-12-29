# docker-exposer

## Description

This is a simple server that exposes the docker socket to the outside world through http.

Although [docker daemon already supports remote access via http](https://docs.docker.com/config/daemon/remote-access/),
built-in authentication method is limited to TLS.

You can modify this project and serve docker socket with additional features like authentication or CORS filter.
