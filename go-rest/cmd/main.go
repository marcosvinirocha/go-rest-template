package main

import (
	"log"
	"net/http"
	"go-rest/api/routes"
	"go-rest/api/config"
	"go-rest/api/middleware"
)

func main() {
	// Load configuration
	cfg := config.LoadConfig()

	// Initialize database connection
	db := config.InitDatabase(cfg.DatabaseURL)
	defer db.Close()

	// Setup routes
	mux := routes.NewRouter()

	// Apply middleware
	handler := middleware.LoggingMiddleware(mux)
	handler = middleware.AuthMiddleware(handler)

	// Start server
	log.Printf("Server starting on port %s", cfg.Port)
	log.Fatal(http.ListenAndServe(":"+cfg.Port, handler))
}