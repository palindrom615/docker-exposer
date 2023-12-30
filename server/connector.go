package server

import (
	"bufio"
	"io"
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

func (c *Connector) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	res, err := c.RoundTrip(req)
	if err != nil {
		log.Errorw("Failed to read response", "error", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	respond(w, res)
}

func respond(w http.ResponseWriter, res *http.Response) {
	w.WriteHeader(res.StatusCode)
	for key, values := range res.Header {
		for _, value := range values {
			w.Header().Add(key, value)
		}
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Errorw("Failed to read response body", "error", err)
		return
	}
	w.Write(body)
}
