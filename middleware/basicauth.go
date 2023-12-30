package middleware

import (
	"net/http"
)

type BasicAuth struct {
	username string
	password string
}

func NewBasicAuth(username, password string) *BasicAuth {
	return &BasicAuth{username: username, password: password}
}

func (b *BasicAuth) Middleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		username, password, ok := r.BasicAuth()
		if !ok || username != b.username || password != b.password {
			w.Header().Set("WWW-Authenticate", `Basic realm="Restricted", charset="UTF-8"`)
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		handler.ServeHTTP(w, r)
	})
}
