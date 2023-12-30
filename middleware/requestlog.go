package middleware

import (
	"context"
	"github.com/palindrom615/docker-exposer/logger"
	"math/rand"
	"net/http"
)

type key string

const RequestIDKey key = "request_id"

var log = logger.DefaultLogger()

func RequestLog(roundTrip http.RoundTripper) http.RoundTripper {
	return RoundTripFunc(
		func(req *http.Request) (res *http.Response, err error) {
			rid := rand.Uint64()
			log := log.With(RequestIDKey, rid)
			req = req.WithContext(context.WithValue(req.Context(), RequestIDKey, rid))
			log.Debugw("Request", "method", req.Method, "url", req.URL, "headers", req.Header, "body", req.Body)

			res, err = roundTrip.RoundTrip(req)

			log.Debugw("Response", "method", req.Method, "url", req.URL, "headers", res.Header, "body", res.Body, "err", err)
			return
		},
	)
}
