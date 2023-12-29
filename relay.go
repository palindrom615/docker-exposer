package docker_exposer

import (
	"bufio"
	"net"
	"net/http"
)

type Relay struct {
	conn net.Conn
}

func NewRelay(conn net.Conn) *Relay {
	return &Relay{conn: conn}
}

func (r *Relay) RoundTrip(req *http.Request) (*http.Response, error) {
	if err := req.Write(r.conn); err != nil {
		return nil, err
	}
	return http.ReadResponse(bufio.NewReader(r.conn), req)
}
