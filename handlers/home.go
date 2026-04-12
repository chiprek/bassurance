package handlers

import (
	"html/template"
	"log"
	"net/http"
	"github.com/chiprek/bassurance/store"
)

func (a *App) Home(w http.ResponseWriter, r *http.Request) {

	processes, err := GetAllProcesses(a.DB)
	if err != nil {
		log.Println("Database error:", err)
		http.Error(w, "Failed to load blueprints", http.StatusInternalServerError)
		return
	}

	tmpl, err := template.ParseFiles("templates/base.html", "templates/home.html")
	if err != nil {
		log.Println("Template parsing error:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	err = tmpl.ExecuteTemplate(w, "base", processes)
	if err != nil {
		log.Println("Template execution error:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
