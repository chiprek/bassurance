package main

import (
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/chiprek/bassurance/internal/database"
	"github.com/google/uuid"
)

func (cfg *apiConfig) handleCreateUnit(w http.ResponseWriter, r *http.Request) {
	jobName := r.PathValue("name")

	jobDBName := normalize(jobName)

	type parameters struct {
		SerialNumber string `json:"serialnumber"`
		Status       string `json:"status"`
	}

	type response struct {
		ID            uuid.UUID `json:"id"`
		Job_id        uuid.UUID `json:"job_id"`
		Serial_number string    `json:"serial_number"`
		Created_at    time.Time `json:"created_at"`
		Status        string    `json:"status"`
	}

	decoder := json.NewDecoder(r.Body)

	requestParams := parameters{}
	err := decoder.Decode(&requestParams)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "something went wrong")
		return
	}

	jobDBID, err := cfg.database.GetJob(r.Context(), jobDBName)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			respondWithError(w, http.StatusNotFound, "Job not found")
			return
		} else {
			respondWithError(w, http.StatusInternalServerError, "something went wrong")
		}
		return
	}

	var dbSerialNumber sql.NullString

	if requestParams.SerialNumber != "" {
		dbSerialNumber = sql.NullString{
			String: requestParams.SerialNumber,
			Valid:  true,
		}
	} else {
		dbSerialNumber = sql.NullString{
			String: "",
			Valid:  false,
		}
	}

	dbParams := database.CreateUnitParams{
		JobID:        jobDBID.ID,
		SerialNumber: dbSerialNumber,
		Status:       requestParams.Status,
	}

	dbUnit, err := cfg.database.CreateUnit(r.Context(), dbParams)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "failed to create unit")
		return
	}

	respondWithJSON(w, http.StatusCreated, response{ID: dbUnit.ID, Job_id: dbUnit.JobID, Serial_number: dbUnit.SerialNumber.String, Created_at: dbUnit.CreatedAt, Status: dbUnit.Status})

}
