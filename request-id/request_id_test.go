package requestid

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRequestIDMiddlewareSetsResponseHeader(t *testing.T) {
	next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {})

	handler := RequestIDMiddleware(next)

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	recorder := httptest.NewRecorder()

	handler.ServeHTTP(recorder, req)

	id := recorder.Header().Get("X-Request-ID")

	if id == "" {
		t.Fatal("expected X-Request-ID header")
	}

	if len(id) != 32 {
		t.Errorf("expected 32-character ID, got %d", len(id))
	}
}

func TestRequestIDMiddlewareStoresIDInContext(t *testing.T) {
	var contextID string

	next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		contextID = GetRequestID(r.Context())
	})

	handler := RequestIDMiddleware(next)

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	recorder := httptest.NewRecorder()

	handler.ServeHTTP(recorder, req)

	headerID := recorder.Header().Get("X-Request-ID")

	if contextID == "" {
		t.Fatal("expected request ID in context")
	}

	if contextID != headerID {
		t.Errorf("context ID %q does not match header ID %q", contextID, headerID)
	}
}

func TestRequestIDMiddlewareGeneratesUniqueIDs(t *testing.T) {
	next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {})

	handler := RequestIDMiddleware(next)

	firstRecorder := httptest.NewRecorder()
	firstReq := httptest.NewRequest(http.MethodGet, "/", nil)
	handler.ServeHTTP(firstRecorder, firstReq)

	secondRecorder := httptest.NewRecorder()
	secondReq := httptest.NewRequest(http.MethodGet, "/", nil)
	handler.ServeHTTP(secondRecorder, secondReq)

	firstID := firstRecorder.Header().Get("X-Request-ID")
	secondID := secondRecorder.Header().Get("X-Request-ID")

	if firstID == secondID {
		t.Fatal("expected different request IDs")
	}
}

func TestGetRequestIDReturnsEmptyWhenMissing(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)

	id := GetRequestID(req.Context())

	if id != "" {
		t.Errorf("expected empty ID, got %q", id)
	}
}
