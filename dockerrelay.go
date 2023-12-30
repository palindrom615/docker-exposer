package docker_exposer

import (
	"bufio"
	"context"
	"github.com/docker/docker/client"
	"github.com/palindrom615/docker-exposer/logger"
	"io"
	"net"
	"net/http"
	"time"
)

var log = logger.DefaultLogger()

type DockerRelay struct {
	dockerClient *client.Client
	conn         net.Conn
}

func NewDockerRelay(opts ...client.Opt) *DockerRelay {
	cli, err := client.NewClientWithOpts(opts...)
	if err != nil {
		panic(err)
	}

	d := &DockerRelay{
		dockerClient: cli,
	}
	d.dialAgain()
	return d
}

func (d *DockerRelay) Close() error {
	return d.dockerClient.Close()
}

func (d *DockerRelay) dialAgain() {
	const connectionRetryWait = 5 * time.Second

	for {
		conn, err := d.dockerClient.Dialer()(context.Background())
		if err != nil {
			log.Errorw("Failed to dial docker; try again", "error", err)
			time.Sleep(connectionRetryWait)
			continue
		}
		log.Infow(
			"Connected to docker. Ready to relay.",
			"addr", conn.RemoteAddr().String(),
			"local", conn.LocalAddr().String(),
			"network", conn.RemoteAddr().Network(),
			"type", conn.RemoteAddr().Network(),
		)
		d.conn = conn
		break
	}
	return
}

func (c *DockerRelay) RoundTrip(req *http.Request) (*http.Response, error) {
	if err := req.Write(c.conn); err != nil {
		log.Errorw("Failed to write request", "error", err)
		go c.dialAgain()

		return nil, err
	}
	return http.ReadResponse(bufio.NewReader(c.conn), req)
}

func (d *DockerRelay) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	res, err := d.RoundTrip(req)
	if err != nil {
		log.Errorw("Failed to read response", "error", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer res.Body.Close()

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
