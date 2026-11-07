package fundamentals

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, Go!")
}

func AboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Learning Go Middleware")
}

func HealthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	response := map[string]string{
		"status": "UP",
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		return
	}
}

func SlowHandler(w http.ResponseWriter, r *http.Request) {
	time.Sleep(2 * time.Second)

	fmt.Fprintln(w, "That was slow!")
}

func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Nothing to see here", http.StatusNotFound)
}

func PanicHandler(w http.ResponseWriter, r *http.Request) {
	panic("intentional test panic")
}
