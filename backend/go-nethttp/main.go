package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

type MessageRequest struct {
	Message string `json:"message"`
}

type MessageResponse struct {
	Message string `json:"message"`
}

func main() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	port := os.Getenv("PORT")
	frontendOrigin := os.Getenv("FRONTEND_ORIGIN")

	mux := http.NewServeMux()

	mux.HandleFunc("GET /api/message", handleGetMessage)
	mux.HandleFunc("POST /api/message", handlePostMessage)

	fmt.Printf("Server running on http://localhost:%s\n", port)

	err = http.ListenAndServe(":"+port, corsMiddleware(mux, frontendOrigin))
	if err != nil {
		panic(err)
	}
}

func handleGetMessage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(MessageResponse{
		Message: "Hello from GO-NETHTTP backend",
	})
}

func handlePostMessage(w http.ResponseWriter, r *http.Request) {
	var body MessageRequest

	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(MessageResponse{
		Message: fmt.Sprintf("Message received by GO NETHTTP backend: %q", body.Message),
	})
}

func corsMiddleware(next http.Handler, frontendOrigin string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", frontendOrigin)
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		next.ServeHTTP(w, r)
	})
}
