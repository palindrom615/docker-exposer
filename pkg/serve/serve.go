package serve

import (
	"github.com/palindrom615/docker-exposer/pkg/logger"
	"io"
	"net/http"
)

var log = logger.DefaultLogger()

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
