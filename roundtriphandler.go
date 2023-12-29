package docker_exposer

import (
	"io"
	"net/http"
)

type RoundTripHandler struct {
	roundTripper http.RoundTripper
}

func NewRoundTripHandler(roundTripper http.RoundTripper) *RoundTripHandler {
	return &RoundTripHandler{roundTripper: roundTripper}
}

func (r *RoundTripHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	log := DefaultLogger()
	res, err := r.roundTripper.RoundTrip(req)
	if err != nil {
		log.Error("Failed to read response", "error", err)
		return
	}
	defer res.Body.Close()
	writeResponse(w, res)
}

func writeResponse(w http.ResponseWriter, res *http.Response) {
	w.WriteHeader(res.StatusCode)
	for key, values := range res.Header {
		for _, value := range values {
			w.Header().Add(key, value)
		}
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		DefaultLogger().Error("Failed to read response body", "error", err)
		return
	}
	w.Write(body)
}
