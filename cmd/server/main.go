package main

import (
	"log"
	"net/http"

	"github.com/Ramadan-Lotfy/go-middleware-mastery/fundamentals"
	"github.com/Ramadan-Lotfy/go-middleware-mastery/logging"
	"github.com/Ramadan-Lotfy/go-middleware-mastery/middleware-chain"
	"github.com/Ramadan-Lotfy/go-middleware-mastery/recovery"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /hello", fundamentals.HelloHandler)
	mux.HandleFunc("GET /about", fundamentals.AboutHandler)
	mux.HandleFunc("GET /health", fundamentals.HealthHandler)
	mux.HandleFunc("GET /slow", fundamentals.SlowHandler)
	mux.HandleFunc("GET /not-found", fundamentals.NotFoundHandler)
	// mux.HandleFunc("GET /panic", fundamentals.PanicHandler)

	handler := middlewarechain.Chain(
		mux,
		logging.LoggingMiddleware,
		recovery.RecoveryMiddleware,
	)

	addr := ":8080"

	log.Printf("Server listening on %s", addr)

	if err := http.ListenAndServe(addr, handler); err != nil {
		log.Fatal(err)
	}
}
