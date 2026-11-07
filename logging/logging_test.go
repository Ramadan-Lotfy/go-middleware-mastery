package logging

import (
	"bytes"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/Ramadan-Lotfy/go-middleware-mastery/middleware-chain"
	"github.com/Ramadan-Lotfy/go-middleware-mastery/request-id"
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

func TestLoggingMiddlewareIncludesRequestID(t *testing.T) {
	var logOutput bytes.Buffer

	originalOutput := log.Writer()
	log.SetOutput(&logOutput)
	defer log.SetOutput(originalOutput)

	next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	handler := middlewarechain.Chain(
		next,
		requestid.RequestIDMiddleware,
		LoggingMiddleware,
	)

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	recorder := httptest.NewRecorder()

	handler.ServeHTTP(recorder, req)

	id := recorder.Header().Get("X-Request-ID")

	if id == "" {
		t.Fatal("expected request ID in response header")
	}

	output := logOutput.String()

	if !strings.Contains(output, id) {
		t.Errorf("expected logs to contain request ID %q, got %q", id, output)
	}
}
