package main

import (
	"log"
	"net/http"
	"os"

	"github.com/dgraph-io/badger/v3"
	"github.com/go-chi/chi/v5"
	"github.com/hculpan/kabbase/pkg/dbbadger"
	"github.com/joho/godotenv"
)

var secretKey []byte
var db *badger.DB

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

	key := os.Getenv("PASETO_SECRET_KEY")
	if len(key) == 0 {
		log.Fatal("Unable to find secret key")
	}
	secretKey = []byte(key)

	// Initialize BadgerDB
	dbConn, err := dbbadger.OpenDB(dbPath)
	if err != nil {
		log.Fatal(err)
	}
	db = dbConn
	defer db.Close()

	// Chi Router
	r := chi.NewRouter()

	// Middleware for PASETO token verification
	r.Use(LogMiddleware)
	r.Use(AuthMiddleware)

	setRoutes(r)

	log.Default().Printf("Data path: %s\n", dbPath)
	log.Default().Printf("Server started on port %s\n", port)
	http.ListenAndServe(":"+port, r)
}
