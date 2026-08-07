package main

import (
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/chiprek/bassurance/internal/database"
	"github.com/google/uuid"
	"github.com/lib/pq"
)

func (cfg *apiConfig) handleCreateUnit(w http.ResponseWriter, r *http.Request) {

	jobDBName := normalize(r.PathValue("name"))

	type parameters struct {
		SerialNumber string `json:"serialnumber"`
	}

	type response struct {
		ID            uuid.UUID `json:"id"`
		Job_id        uuid.UUID `json:"job_id"`
		Serial_number string    `json:"serial_number"`
		Created_at    time.Time `json:"created_at"`
	}

	decoder := json.NewDecoder(r.Body)

	requestParams := parameters{}
	err := decoder.Decode(&requestParams)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "something went wrong")
		return
	}

	jobDB, err := cfg.Queries.GetJob(r.Context(), jobDBName)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			respondWithError(w, http.StatusNotFound, "Job not found")
			return
		} else {
			respondWithError(w, http.StatusInternalServerError, "something went wrong")
		}
		return
	}
	unitID := uuid.New()

	tx, err := cfg.DB.BeginTx(r.Context(), nil)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Something went wrong")
		return
	}
	defer tx.Rollback()

	qtx := cfg.Queries.WithTx(tx)

	unitparams := database.CreateUnitParams{
		ID:           unitID,
		SerialNumber: requestParams.SerialNumber,
	}

	unit, err := qtx.CreateUnit(r.Context(), unitparams)
	if err != nil {

		if pqErr, ok := err.(*pq.Error); ok {
			if pqErr.Code == "23505" {
				respondWithError(w, http.StatusConflict, "A unit with this serial number already exists")
				return
			}
		}

		respondWithError(w, http.StatusInternalServerError, "Failed to create unit")
		return
	}

	jobUnitParams := database.CreateJobUnitParams{
		JobID:  jobDB.ID,
		UnitID: unitID,
	}

	err = qtx.CreateJobUnit(r.Context(), jobUnitParams)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to create link between job and unit")
		return
	}

	err = tx.Commit()

	respondWithJSON(w, http.StatusCreated, response{ID: unit.ID, Job_id: jobDB.ID, Serial_number: unit.SerialNumber, Created_at: unit.CreatedAt.Time})

}

func (cfg *apiConfig) handleGetUnitsByJob(w http.ResponseWriter, r *http.Request) {

	jobDBName := normalize(r.PathValue("name"))

	limit := int32(50)
	if limitStr := r.URL.Query().Get("limit"); limitStr != "" {
		if parsedLimit, err := strconv.Atoi(limitStr); err == nil {
			limit = int32(parsedLimit)
		}
	}

	offset := int32(0)
	if offsetStr := r.URL.Query().Get("offset"); offsetStr != "" {
		if parsedOffset, err := strconv.Atoi(offsetStr); err == nil {
			offset = int32(parsedOffset)
		}
	}

	type response struct {
		ID            uuid.UUID `json:"id"`
		Serial_number string    `json:"serial_number"`
		CreatedAt     time.Time `json:"created_at"`
	}

	jobNameParam := database.GetUnitsByJobNameParams{Name: jobDBName, Limit: limit, Offset: offset}
	dbUnits, err := cfg.Queries.GetUnitsByJobName(r.Context(), jobNameParam)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			respondWithError(w, http.StatusNotFound, "Job not found")
			return
		}

		respondWithError(w, http.StatusInternalServerError, "Failed to retrive units")
		return
	}

	var responseSlice []response
	for _, dbUnit := range dbUnits {
		responseSlice = append(responseSlice, response{
			ID:            dbUnit.ID,
			Serial_number: dbUnit.SerialNumber,
			CreatedAt:     dbUnit.CreatedAt.Time,
		})
	}

	respondWithJSON(w, http.StatusOK, responseSlice)

}

func (cfg *apiConfig) handlerAttachUnits(w http.ResponseWriter, r *http.Request) {
	jobdbname := r.PathValue("name")

	type parameters struct {
		SerialNumber string `json:"serialnumber"`
	}

	dc := json.NewDecoder(r.Body)

	prams := parameters{}

	err := dc.Decode(&prams)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "something went wrong")
		return
	}

	dbj, err := cfg.Queries.GetJob(r.Context(), jobdbname)
	{
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				respondWithError(w, http.StatusNotFound, "Job not found")
				return
			} else {
				respondWithError(w, http.StatusInternalServerError, "something went wrong")
			}
			return
		}
	}

	dbu, err := cfg.Queries.GetUnitBySerialNumber(r.Context(), prams.SerialNumber)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			respondWithError(w, http.StatusNotFound, "Job not found")
			return
		} else {
			respondWithError(w, http.StatusInternalServerError, "something went wrong")
		}
		return
	}

	juparams := database.CreateJobUnitParams{
		JobID:  dbj.ID,
		UnitID: dbu,
	}

	err = cfg.Queries.CreateJobUnit(r.Context(), juparams)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "something went wrong")
		return
	}
	w.WriteHeader(http.StatusNoContent)

}
