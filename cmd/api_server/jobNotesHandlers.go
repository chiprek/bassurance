package main

import (
	"encoding/json"
	"net/http"
)

func (cfg *apiConfig) HandleCreateJobNote(w http.ResponseWriter, r *http.Request) {
	UrlName := r.PathValue("name")

	sanitized := normalize(UrlName)

	type request struct {
		subAsmbId string
	}

	decoder := json.NewDecoder(r.Body)
	params := request{}

	respondWithJSON()
}
