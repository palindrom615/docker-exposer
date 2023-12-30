package server

import (
	"bufio"
	"net"
	"net/http"
)

type Connector struct {
	conn net.Conn
}

func NewConnector(conn net.Conn) *Connector {
	return &Connector{conn: conn}
}

func (c *Connector) RoundTrip(req *http.Request) (*http.Response, error) {
	if err := req.Write(c.conn); err != nil {
		return nil, err
	}
	return http.ReadResponse(bufio.NewReader(c.conn), req)
}
