package main

import (
	"database/sql"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"sort"
	"time"

	"github.com/chiprek/bassurance/internal/database"
	"github.com/google/uuid"
)

// handleCreateJob Creates a Job by passing it 2 parameters a NAME and an optional Status
// at this point in development it will return the generated uuid, creation time, name, and status if given.
func (cfg *apiConfig) handleCreateJob(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name   string `json:"name"`
		Status string `json:"status"`
	}
	type response struct {
		ID         uuid.UUID `json:"id"`
		Created_at time.Time `json:"created_at"`
		Name       string    `json:"name"`
		Status     string    `json:"status,omitempty"`
	}
	decoder := json.NewDecoder(r.Body)

	requestParams := parameters{}
	err := decoder.Decode(&requestParams)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "something went wrong")
		return
	}

	requestParams.Name = normalize(requestParams.Name)

	newJobID := uuid.New()

	params := database.CreateJobParams{
		ID:     newJobID,
		Name:   requestParams.Name,
		Status: requestParams.Status,
	}

	createJob, err := cfg.database.CreateJob(r.Context(), params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "not able to create job")
		return
	}
	respondWithJSON(w, http.StatusCreated, response{ID: createJob.ID, Created_at: createJob.CreatedAt.Time, Name: createJob.Name, Status: createJob.Status})
}

func (cfg *apiConfig) handlerGetJobs(w http.ResponseWriter, r *http.Request) {
	sortDirection := r.URL.Query().Get("sort")

	var dbJobs []database.Job
	var err error

	dbJobs, err = cfg.database.GetJobs(r.Context())
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "something went wrong")
		log.Printf("error: %v", err)
		return
	}

	allJobs := make([]Jobs, 0, len(dbJobs))

	for _, dbJobs := range dbJobs {
		allJobs = append(allJobs, Jobs{
			ID:         dbJobs.ID,
			Name:       dbJobs.Name,
			Status:     dbJobs.Status,
			Created_at: dbJobs.CreatedAt.Time,
			Updated_at: dbJobs.UpdatedAt.Time,
		})
	}

	switch sortDirection {
	case "desc":
		sort.Slice(allJobs, func(i, j int) bool {
			return allJobs[i].Created_at.After(allJobs[j].Created_at)
		})
	case "asc", "":
		fallthrough
	default:
		sort.Slice(allJobs, func(i, j int) bool {
			return allJobs[i].Created_at.Before(allJobs[j].Created_at)
		})
	}
	respondWithJSON(w, http.StatusOK, allJobs)
}

func (cfg *apiConfig) handlerGetSpecifiedJob(w http.ResponseWriter, r *http.Request) {
	UrlName := r.PathValue("name")

	sanitized := normalize(UrlName)

	dbJob, err := cfg.database.GetJob(r.Context(), sanitized)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			respondWithError(w, http.StatusNotFound, "Job not found")
			return
		} else {
			respondWithError(w, http.StatusInternalServerError, "something went wrong")
		}
		return
	}

	respondWithJSON(w, http.StatusOK, Jobs{ID: dbJob.ID, Created_at: dbJob.CreatedAt.Time, Updated_at: dbJob.UpdatedAt.Time, Name: dbJob.Name, Status: dbJob.Status})

}
