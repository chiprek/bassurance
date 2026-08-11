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
	Platform string
	DB       *sql.DB
	Queries  *database.Queries
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
		Platform: platform,
		DB:       db,
		Queries:  dbQueries,
	}
	mux := http.NewServeMux()
	mux.HandleFunc("POST /api/v1/jobs", cfg.handleCreateJob)
	mux.HandleFunc("GET /api/v1/jobs", cfg.handlerGetJobs)
	mux.HandleFunc("GET /api/v1/jobs/{name}", cfg.handlerGetSpecifiedJob)
	mux.HandleFunc("POST /api/v1/jobs/{name}/units", cfg.handleCreateUnit)
	mux.HandleFunc("GET /api/v1/jobs/{name}/units", cfg.handleGetUnitsByJob)
	mux.HandleFunc("POST /api/v1/jobs/{name}/units/attach", cfg.handlerAttachUnits)
	mux.HandleFunc("POST /api/v1/units/{serial_number}/sub-assemblies", cfg.handleCreateSubAssembly)
	mux.HandleFunc("GET /api/v1/units/{serial_number}/sub-assemblies", cfg.handleGetSubAssemblies)
	server := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	log.Printf("Server Listening on port %s...\n", port)
	log.Fatal(server.ListenAndServe())
}
