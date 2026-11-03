package fundamentals

import (
    "net/http"
    "net/http/httptest"
    "testing"
)

func TestHelloHandler(t *testing.T) {
    req := httptest.NewRequest(http.MethodGet, "/hello", nil)
    recorder := httptest.NewRecorder()

    HelloHandler(recorder, req)

    if recorder.Code != http.StatusOK {
        t.Errorf("expected status %d, got %d",
            http.StatusOK, recorder.Code)
    }

    expected := "Hello, Go!\n"

    if recorder.Body.String() != expected {
        t.Errorf("expected %q, got %q",
            expected, recorder.Body.String())
    }
}

func TestHealthHandler(t *testing.T) {
    req := httptest.NewRequest(http.MethodGet, "/health", nil)
    recorder := httptest.NewRecorder()

    HealthHandler(recorder, req)

    if recorder.Code != http.StatusOK {
        t.Errorf("expected status %d, got %d",
            http.StatusOK, recorder.Code)
    }

    if recorder.Header().Get("Content-Type") != "application/json" {
        t.Error("expected application/json content type")
    }
}


func TestAboutHandler(t *testing.T) {
    req := httptest.NewRequest(http.MethodGet, "/about", nil)
    recorder := httptest.NewRecorder()

    AboutHandler(recorder, req)

    if recorder.Code != http.StatusOK {
        t.Errorf("expected status 200, got %d", recorder.Code)
    }

    expected := "Learning Go Middleware\n"

    if recorder.Body.String() != expected {
        t.Errorf("expected %q, got %q",
            expected, recorder.Body.String())
    }
}

func TestNotFoundHandler(t *testing.T) {
    req := httptest.NewRequest(http.MethodGet, "/not-found", nil)
    recorder := httptest.NewRecorder()

    NotFoundHandler(recorder, req)

    if recorder.Code != http.StatusNotFound {
        t.Errorf("expected status 404, got %d", recorder.Code)
    }
}