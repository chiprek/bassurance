package main

import (
	"net/http"
	"os"
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
	defer file.Close()

	var jobName string
	var subAssembly string

	distinctDir := filepath.Join("./uploads", jobName, subAssembly)

	filename := filepath.Base(header.Filename)
	fullPath := filepath.Join(distinctDir, filename)

	dst, err := os.Create(fullPath)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to save file")
		return
	}

}
