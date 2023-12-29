package middleware

import (
	"io"
	"net/http"
)

type Handler struct {
	middlewares []Middleware
	core        http.RoundTripper
}

func NewHandler(core http.RoundTripper, middlewares ...Middleware) *Handler {
	return &Handler{core: core, middlewares: middlewares}
}

func (h *Handler) Use(middleware ...Middleware) {
	h.middlewares = append(h.middlewares, middleware...)
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	r := chain(h.middlewares...)(h.core)
	res, err := r.RoundTrip(req)
	if err != nil {
		log.Errorw("Failed to read response", "error", err)
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
