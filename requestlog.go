package docker_exposer

import (
	"context"
	"log/slog"
	"math/rand"
	"net/http"
)

var RequestIDKey = "request_id"

type RequestLog struct {
	log  *slog.Logger
	next http.RoundTripper
}

func NewRequestLog(log *slog.Logger, next http.RoundTripper) *RequestLog {
	return &RequestLog{log: log, next: next}
}

func (r *RequestLog) RoundTrip(req *http.Request) (res *http.Response, err error) {
	rid := rand.Uint64()
	log := r.log.With(RequestIDKey, rid)
	req = req.WithContext(context.WithValue(req.Context(), RequestIDKey, rid))
	log.Debug("Request", "req", req)
	res, err = r.next.RoundTrip(req)
	log.Debug("Response", "res", res, "err", err)
	return
}
