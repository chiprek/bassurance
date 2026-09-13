package main

import "net/http"

func (cfg *apiConfig) HandleCreateJobNote(w http.ResponseWriter, r *http.Request) {
	UrlName := r.PathValue("name")

	sanitized := normalize(UrlName)

	type request struct {
		subAsmbId string
	}

	respondWithJSON()
}
