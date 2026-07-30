package main

import (
	"encoding/json"
	"net/http"
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
	params := database.CreateJobParams{
		Name:   requestParams.Name,
		Status: requestParams.Status,
	}

	created, err := cfg.database.CreateJob(r.Context(), params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "not able to create job")
		return
	}
	respondWithJSON(w, http.StatusCreated, response{ID: created.ID, Created_at: created.CreatedAt, Name: created.Name, Status: created.Status})
}
