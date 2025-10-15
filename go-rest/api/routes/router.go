package routes

import (
	"net/http"
	"go-rest/api/handlers"
)

func NewRouter() http.Handler {
	mux := http.NewServeMux()

	// Define routes
	mux.HandleFunc("GET /users", handlers.GetAllUsers)
	mux.HandleFunc("GET /users/{id}", handlers.GetUserByID)
	mux.HandleFunc("POST /users", handlers.CreateUser)
	mux.HandleFunc("PUT /users/{id}", handlers.UpdateUser)
	mux.HandleFunc("DELETE /users/{id}", handlers.DeleteUser)

	// Health check endpoint
	mux.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status": "healthy"}`))
	})

	return mux
}