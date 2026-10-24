package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "Invaliad request - not found",
		})
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Hello from my Go DevOps API!",
	})
}

func versionHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"version": "v3",
	})
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"status": "healthy",
	})
}

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/health", healthHandler)
	http.HandleFunc("/version", versionHandler)

	log.Println("API server starting on :8080")

	log.Fatal(http.ListenAndServe(":8080", nil))
}
