package middleware

import (
	"net/http"
)

type Middleware func(roundTrip http.RoundTripper) http.RoundTripper

type RoundTripFunc func(req *http.Request) (res *http.Response, err error)

func (f RoundTripFunc) RoundTrip(req *http.Request) (res *http.Response, err error) {
	return f(req)
}

func chain(middleware ...Middleware) Middleware {
	return func(roundTrip http.RoundTripper) http.RoundTripper {
		for _, m := range middleware {
			roundTrip = m(roundTrip)
		}
		return roundTrip
	}
}
