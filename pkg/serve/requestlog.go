package serve

import (
	"context"
	"math/rand"
	"net/http"
)

type key string

const RequestIDKey key = "request_id"

type RequestLog struct {
	next http.RoundTripper
}

func NewRequestLog(next http.RoundTripper) *RequestLog {
	return &RequestLog{next: next}
}

func (r *RequestLog) RoundTrip(req *http.Request) (res *http.Response, err error) {
	rid := rand.Uint64()
	log := log.With(RequestIDKey, rid)
	req = req.WithContext(context.WithValue(req.Context(), RequestIDKey, rid))
	log.Debugw("Request", "method", req.Method, "url", req.URL, "headers", req.Header, "body", req.Body)
	res, err = r.next.RoundTrip(req)
	log.Debugw("Response", "method", req.Method, "url", req.URL, "headers", res.Header, "body", res.Body, "err", err)
	return
}
