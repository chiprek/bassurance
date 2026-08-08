package main

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/chiprek/bassurance/internal/database"
	"github.com/google/uuid"
)

func (cfg *apiConfig) handleCreateSubAssembly(w http.ResponseWriter, r *http.Request) {
	unitSN := r.PathValue("serial_number")

	type request struct {
		Name         string `json:"name"`
		SerialNumber string `json:"serial_number"`
		Status       string `json:"status"`
	}

	decoder := json.NewDecoder(r.Body)
	params := request{}

	if err := decoder.Decode(&params); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid JSON playload")
		return
	}
	unitID, err := cfg.Queries.GetUnitBySerialNumber(r.Context(), unitSN)
	if err != nil {
		respondWithError(w, http.StatusNotFound, "Parent unit not found")
		return
	}

	dbSN := sql.NullString{
		String: params.SerialNumber,
		Valid:  params.SerialNumber != "",
	}

	subAssembly, err := cfg.Queries.CreateSubAssembly(r.Context(), database.CreateSubAssemblyParams{
		ID:           uuid.New(),
		UnitID:       unitID,
		Name:         params.Name,
		SerialNumber: dbSN,
		Status:       params.Status,
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "failed to make sub-assembly")
		return
	}

	respondWithJSON(w, http.StatusOK, subAssembly)
}
