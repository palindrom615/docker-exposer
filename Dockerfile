# syntax=docker/dockerfile:1
FROM golang:1.21

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . ./

RUN CGO_ENABLED=0 GOOS=linux go build ./cmd/docker-exposer

EXPOSE 8080

CMD ["./docker-exposer"]
