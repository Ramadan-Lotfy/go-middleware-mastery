package logging

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestLoggingMiddlewareCallsNextHandler(t *testing.T) {
	called := false

	next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		called = true
		w.WriteHeader(http.StatusOK)
	})

	middleware := LoggingMiddleware(next)

	req := httptest.NewRequest(http.MethodGet, "/hello", nil)
	recorder := httptest.NewRecorder()

	middleware.ServeHTTP(recorder, req)

	if !called {
		t.Fatal("expected next handler to be called")
	}

	if recorder.Code != http.StatusOK {
		t.Errorf("expected status 200, got %d", recorder.Code)
	}
}

func TestLoggingMiddlewareCapturesStatusCode(t *testing.T) {
	next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Not found", http.StatusNotFound)
	})

	middleware := LoggingMiddleware(next)

	req := httptest.NewRequest(http.MethodGet, "/missing", nil)
	recorder := httptest.NewRecorder()

	middleware.ServeHTTP(recorder, req)

	if recorder.Code != http.StatusNotFound {
		t.Errorf("expected status 404, got %d", recorder.Code)
	}
}
