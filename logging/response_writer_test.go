package logging

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestResponseWriterDefaultsToOK(t *testing.T) {
	recorder := httptest.NewRecorder()
	rw := newResponseWriter(recorder)

	_, err := rw.Write([]byte("Hello"))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if rw.statusCode != http.StatusOK {
		t.Errorf("expected status 200, got %d", rw.statusCode)
	}

	if recorder.Body.String() != "Hello" {
		t.Errorf("expected body Hello, got %q", recorder.Body.String())
	}
}

func TestResponseWriterCapturesExplicitStatus(t *testing.T) {
	recorder := httptest.NewRecorder()
	rw := newResponseWriter(recorder)

	rw.WriteHeader(http.StatusNotFound)

	if rw.statusCode != http.StatusNotFound {
		t.Errorf("expected status 404, got %d", rw.statusCode)
	}

	if recorder.Code != http.StatusNotFound {
		t.Errorf("expected recorder status 404, got %d", recorder.Code)
	}
}

func TestResponseWriterKeepsFirstStatus(t *testing.T) {
	recorder := httptest.NewRecorder()
	rw := newResponseWriter(recorder)

	rw.WriteHeader(http.StatusCreated)
	rw.WriteHeader(http.StatusNotFound)

	if rw.statusCode != http.StatusCreated {
		t.Errorf("expected status 201, got %d", rw.statusCode)
	}

	if recorder.Code != http.StatusCreated {
		t.Errorf("expected recorder status 201, got %d", recorder.Code)
	}
}
