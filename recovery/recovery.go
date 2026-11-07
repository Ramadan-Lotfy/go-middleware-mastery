package recovery

import (
	"log"
	"net/http"
)

func RecoveryMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// TODO 1: Register a deferred function.
		defer func() {

			// TODO 2: Inside it, call recover().
			if err := recover(); err != nil {

				// TODO 3: If a panic was recovered:
				// Log the panic.
				log.Printf("Recovered from panic: %v", err)

				// Return HTTP 500 using http.Error().
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			}
		}()

		// TODO 4: Call the next handler.
		next.ServeHTTP(w, r)
	})
}
