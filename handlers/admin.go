package handlers

import (
	"html/template"
	"log"
	"net/http"

	"github.com/chiprek/bassurance/store"
)

func (a *App) Admin(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodPost {
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Failed to parse form data", http.StatusBadRequest)
			return

		}

		newProcess := store.Process{
			Name:        r.FormValue("name"),
			Description: r.FormValue("description"),
			CreatedBy:   1,
		}

		var emptySteps []store.Step
		_, err = store.CreateProcessWithSteps(a.DB, newProcess, emptySteps)
		if err != nil {
			log.Println("Database insertion error:", err)
			http.Error(w, "failed to save blueprint to database", http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	tmpl, err := template.ParseFiles("templates/base.html", "templates/admin.html")
	if err != nil {
		log.Println("Template parsing error:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	err = tmpl.ExecuteTemplate(w, "base", nil)
	if err != nil {
		log.Println("Template execution error:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}

}
