package main

import (
	"net/http"
	"path/filepath"
)

func (cfg *apiConfig) HandleCreateJobNote(w http.ResponseWriter, r *http.Request) {
	//parse
	if err := r.ParseMultipartForm(30 << 10); err != nil {
		respondWithError(w, http.StatusBadRequest, "failed to parse form")
		return
	}
	//extract
	file, header, err := r.FormFile("photo")
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "failed to get photo")
		return
	}
	filename := filepath.Base(header.Filename)
	fullPath := filepath.Join()

}
