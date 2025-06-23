package main

import (
	"encoding/json"
	"log"
	"net/http"
)

// Middleware to log each request
func loggingMiddleware(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Request: %s %s from %s", r.Method, r.URL.Path, r.RemoteAddr)
		h(w, r)
	}
}

func main() {
	http.HandleFunc("/", loggingMiddleware(func(w http.ResponseWriter, r *http.Request) {
		jsonResponse(w, map[string]string{
			"message": "Welcome to Service 1",
		})
	}))

	http.HandleFunc("/ping", loggingMiddleware(func(w http.ResponseWriter, r *http.Request) {
		jsonResponse(w, map[string]string{
			"status":  "ok",
			"service": "1",
		})
	}))

	http.HandleFunc("/hello", loggingMiddleware(func(w http.ResponseWriter, r *http.Request) {
		jsonResponse(w, map[string]string{
			"message": "Hello from Service 1",
		})
	}))

	log.Println("Service 1 listening on port 8001...")
	if err := http.ListenAndServe(":8001", nil); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}

func jsonResponse(w http.ResponseWriter, data map[string]string) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
