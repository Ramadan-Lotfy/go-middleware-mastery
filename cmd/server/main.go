package main

import (
    "log"
    "net/http"

    "github.com/Ramadan-Lotfy/go-middleware-mastery/fundamentals"
)

func main() {
    mux := http.NewServeMux()

    mux.HandleFunc("GET /hello", fundamentals.HelloHandler)
    mux.HandleFunc("GET /about", fundamentals.AboutHandler)
    mux.HandleFunc("GET /health", fundamentals.HealthHandler)
    mux.HandleFunc("GET /slow", fundamentals.SlowHandler)
    mux.HandleFunc("GET /not-found", fundamentals.NotFoundHandler)

    addr := ":8080"

    log.Printf("Server listening on %s", addr)

    if err := http.ListenAndServe(addr, mux); err != nil {
        log.Fatal(err)
    }
}