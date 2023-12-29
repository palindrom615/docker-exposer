package serve

import (
	"net/http"
)

type RoundTripHandler struct {
	roundTripper http.RoundTripper
}

func NewRoundTripHandler(roundTripper http.RoundTripper) *RoundTripHandler {
	return &RoundTripHandler{roundTripper: roundTripper}
}

func (r *RoundTripHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	res, err := r.roundTripper.RoundTrip(req)
	if err != nil {
		log.Errorw("Failed to read response", "error", err)
		return
	}
	defer res.Body.Close()
	respond(w, res)
}
