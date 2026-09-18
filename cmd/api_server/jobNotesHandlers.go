package main

import (
	"encoding/json"
	"net/http"

	"github.com/chiprek/bassurance/internal/database"
)

func (cfg *apiConfig) HandleCreateJobNote(w http.ResponseWriter, r *http.Request) {
	UrlID := r.PathValue("id")

	sanitized := normalize(UrlID)

	type request struct {
		subAsmbId string
	}

	decoder := json.NewDecoder(r.Body)
	params := request{}

	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "something went wrong")
		return
	}

	subparams := database.GetSubAssembliesParams{UnitID: sanitized}

	subassembly, err := cfg.Queries.GetSubAssemblies()

	r.ParseMultipartForm(32 << 10)
	respondWithJSON()
}
