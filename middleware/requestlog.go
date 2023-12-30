package middleware

import (
	"context"
	"github.com/palindrom615/docker-exposer/logger"
	"go.uber.org/zap"
	"math/rand"
	"net/http"
)

const RequestIDKey string = "request_id"

type CtxKey string

const RequestID CtxKey = "request_id"

var log = logger.DefaultLogger()

type requestLogWriter struct {
	rw         http.ResponseWriter
	statusCode int
	body       []byte
}

func newRequestLogWriter(rw http.ResponseWriter) *requestLogWriter {
	return &requestLogWriter{rw: rw}
}

func (w *requestLogWriter) Header() http.Header {
	return w.rw.Header()
}

func (w *requestLogWriter) Write(b []byte) (int, error) {
	w.body = b
	return w.rw.Write(b)
}

func (w *requestLogWriter) WriteHeader(statusCode int) {
	w.statusCode = statusCode
	w.rw.WriteHeader(statusCode)
}

func (w *requestLogWriter) Log(log *zap.SugaredLogger) {
	log.Debugw("Response", "status", w.statusCode, "body", string(w.body), "headers", w.Header())
}

func LogRequst(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rid := rand.Uint64()
		log := log.With(RequestIDKey, rid)
		r = r.WithContext(context.WithValue(r.Context(), RequestID, rid))
		log.Debugw("Request", "method", r.Method, "url", r.URL, "headers", r.Header, "body", r.Body)

		lrw := newRequestLogWriter(w)

		next.ServeHTTP(lrw, r)
		lrw.Log(log)

		log.Debugw("Response", "method", r.Method, "url", r.URL, "headers", r.Header, "body", r.Body)
	})
}
