package logging

import (
	"github.com/Ramadan-Lotfy/go-middleware-mastery/request-id"
	"log"
	"net/http"
	"time"
)

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// Read the ID once and reuse it for both log lines.
		id := requestid.GetRequestID(r.Context())

		log.Printf("[%s] --> %s %s", id, r.Method, r.URL.Path)

		rw := newResponseWriter(w)
		next.ServeHTTP(rw, r)

		duration := time.Since(start)

		log.Printf("[%s] <-- %s %s status=%d duration=%v",
			id, r.Method, r.URL.Path, rw.statusCode, duration)
	})
}
