package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/chiprek/bassurance/internal/database"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

const filepathroot = "."
const port = "8080"

type apiConfig struct {
	platform string
	database *database.Queries
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Printf("no .env found")
	}
	platform := os.Getenv("PLATFORM")
	dbURL := os.Getenv("DB_URL")
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal(err)
	}

	dbQueries := database.New(db)

	cfg := &apiConfig{
		platform: platform,
		database: dbQueries,
	}
	mux := http.NewServeMux()
	mux.HandleFunc("POST /api/jobs", cfg.handleCreateJob)

	server := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	log.Printf("Server Listening on port %s...\n", port)
	log.Fatal(server.ListenAndServe())
}
