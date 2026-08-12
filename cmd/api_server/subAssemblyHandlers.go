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

func (cfg *apiConfig) handleGetSubAssemblies(w http.ResponseWriter, r *http.Request) {

	type SubAsm struct {
		ID           uuid.UUID
		Name         string
		SerialNumber *string
		Status       string
		CreatedAt    time.Time
	}

	limit := int32(10)
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

	sortDirection := r.URL.Query().Get("sort")

	unitSN := r.PathValue("serial_number")

	dbUnit, err := cfg.Queries.GetUnitBySerialNumber(r.Context(), unitSN)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			respondWithError(w, http.StatusNotFound, "unit not found")
			return
		}

		respondWithError(w, http.StatusInternalServerError, "unable to fetch unit")
		return
	}

	getSubAsmParams := database.GetSubAssembliesParams{
		UnitID:        dbUnit,
		SortDirection: sortDirection,
		Limit:         limit,
		Offset:        offset,
	}

	var uSubAsm []database.SubAssembly
	uSubAsm, err = cfg.Queries.GetSubAssemblies(r.Context(), getSubAsmParams)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "something went wrong")
		return
	}

	allSubs := make([]SubAsm, 0, len(uSubAsm))
	for _, dbSub := range uSubAsm {

		var sn *string

		if dbSub.SerialNumber.Valid {
			sn = &dbSub.SerialNumber.String
		}

		allSubs = append(allSubs, SubAsm{
			ID:           dbSub.ID,
			Name:         dbSub.Name,
			Status:       dbSub.Status,
			SerialNumber: sn,
			CreatedAt:    dbSub.CreatedAt,
		})
	}

	respondWithJSON(w, http.StatusOK, allSubs)

}
