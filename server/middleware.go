package server

import (
	"log/slog"
	"net/http"
)

func TestRouteHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}

func LoggingMiddleware(logger *slog.Logger, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger.Info("Request", "method", r.Method, "path", r.URL.Path)
		next.ServeHTTP(w, r)
	})
}
