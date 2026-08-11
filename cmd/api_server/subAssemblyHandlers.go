package main

import (
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"
	"sort"

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
	sortDirection := r.URL.Query().Get("sort")

	type SubAssembly struct {
		ID           uuid.UUID
		Name         string
		SerialNumber sql.NullString
		Status       string
		CreatedAt    sql.NullTime
	}

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
	var uSubAsm []database.SubAssembly
	uSubAsm, err = cfg.Queries.GetSubAssemblies(r.Context(), dbUnit)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "something went wrong")
		return
	}

	allSubs := make([]SubAssembly, 0, len(uSubAsm))
	for _, allSub := range allSubs {
		allSubs = append(allSubs, SubAssembly{
			ID:           allSub.ID,
			Name:         allSub.Name,
			Status:       allSub.Status,
			SerialNumber: allSub.SerialNumber,
			CreatedAt:    allSub.CreatedAt,
		})
	}

	switch sortDirection {
	case "desc":
		sort.Slice(allSubs, func(i, j int) bool {
			return allSubs[i].CreatedAt.Time.After(allSubs[j].CreatedAt.Time)
		})
	case "asc", "":
		fallthrough
	default:
		sort.Slice(allSubs, func(i, j int) bool {
			return allSubs[i].CreatedAt.Time.Before(allSubs[j].CreatedAt.Time)
		})
	}

	respondWithJSON(w, http.StatusOK, allSubs)

}
