package requestid

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"net/http"
)

type contextKey struct{}

func newID() (string, error) {
	b := make([]byte, 16)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	return hex.EncodeToString(b), nil
}

func RequestIDMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// TODO 1: Generate a unique request ID.
		id, err := newID()
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		// TODO 2: Set the X-Request-ID response header.
		w.Header().Set("X-Request-ID", id)

		// TODO 3: Store the request ID in the context.
		ctx := context.WithValue(r.Context(), contextKey{}, id)

		// TODO 4: Create a new request containing that context.
		r = r.WithContext(ctx)

		// TODO 5: Pass the updated request to the next handler.
		next.ServeHTTP(w, r)
	})
}

func GetRequestID(ctx context.Context) string {
	// TODO: Retrieve the request ID from the context.
	id, ok := ctx.Value(contextKey{}).(string)
	if !ok {
		return ""
	}
	return id
}
