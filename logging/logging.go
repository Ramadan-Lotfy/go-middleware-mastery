package logging

import (
	"log"
	"net/http"
	"time"
)

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		rw := newResponseWriter(w)

		log.Printf("--> %s %s", r.Method, r.URL.Path)

		next.ServeHTTP(rw, r)

		duration := time.Since(start)

		log.Printf(
			"<-- %s %s status=%d duration=%v",
			r.Method,
			r.URL.Path,
			rw.statusCode,
			duration,
		)
	})
}
