package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	dbbadger "github.com/hculpan/kabbase/pkg/dbBadger"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbPath := os.Getenv("DB_PATH")
	if len(dbPath) == 0 {
		dbPath = "./data"
	}

	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8080"
	}

	// Initialize BadgerDB
	db, err := dbbadger.OpenDB(dbPath)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Chi Router
	r := chi.NewRouter()

	// Middleware for PASETO token verification
	r.Use(AuthMiddleware)

	// Authentication Endpoint
	r.Post("/authenticate", func(w http.ResponseWriter, r *http.Request) {
		// Implement authentication logic and return PASETO token
	})

	// Protected Endpoints
	r.Get("/protected-resource", func(w http.ResponseWriter, r *http.Request) {
		// This endpoint is protected by the AuthMiddleware
	})

	log.Default().Printf("Data path: %s\n", dbPath)
	log.Default().Printf("Server started on port %s\n", port)
	http.ListenAndServe(":"+port, r)
}

// AuthMiddleware to validate PASETO token
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Implement PASETO token validation logic
		// If valid, call next.ServeHTTP(w, r)
		// If not valid, return an error response
	})
}

// Other functions for handling authentication, token generation, and BadgerDB interactions
