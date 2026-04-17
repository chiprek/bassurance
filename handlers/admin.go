package handlers

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/chiprek/bassurance/store"
)

// logic for the admin section of bassurance
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
		processID, err := store.CreateProcessWithSteps(a.DB, newProcess, emptySteps)
		if err != nil {
			log.Println("Database insertion error:", err)
			http.Error(w, "failed to save blueprint to database", http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, fmt.Sprintf("/admin/builder?id=%d", processID), http.StatusSeeOther)
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

// ProcessBuilder handles the iterative process of adding steps to a blueprint
func (a *App) ProcessBuilder(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	processID, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid bluprint ID", http.StatusBadRequest)
		return
	}

	if r.Method == http.MethodPost {
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Failed to parse form", http.StatusBadRequest)
			return
		}

		action := r.FormValue("action")

		if action == "add_step" {

			process, _ := store.GetProcess(a.DB, processID)

			newStep := store.Step{
				ProcessID:   processID,
				Name:        r.FormValue("name"),
				Description: r.FormValue("description"),
				Required:    r.FormValue("required") != "",
				Critical:    r.FormValue("critical") != "",
				Order:       len(process.Steps) + 1, //auto sequence:
			}

			err = store.AddStep(a.DB, newStep)
			if err != nil {
				log.Println("Database error:", err)
				http.Error(w, "Failed to add step", http.StatusInternalServerError)
				return
			}

		}

		if action == "add_field" {
			stepID, _ := strconv.Atoi(r.FormValue("step_id"))

			var targetVal, tolerance float64
			if t := r.FormValue("target_val"); t != "" {
				targetVal, _ = strconv.ParseFloat(t, 64)
			}
			if tol := r.FormValue("tolerance"); tol != "" {
				tolerance, _ = strconv.ParseFloat(tol, 64)
			}

			newField := store.StepField{
				StepID:      stepID,
				Prompt:      r.FormValue("prompt"),
				FieldType:   r.FormValue("field_type"),
				TargetedVal: targetVal,
				Tolerance:   tolerance,
				Order:       1,
			}

			err = store.AddStepField(a.DB, newField)
			if err != nil {
				log.Println("Database error adding field:", err)
				http.Error(w, "Failed to add field", http.StatusInternalServerError)
				return
			}
		}
		http.Redirect(w, r, fmt.Sprintf("/admin/builder?id=%d", processID), http.StatusSeeOther)
		return
	}

	process, err := store.GetProcess(a.DB, processID)
	if err != nil {
		log.Println("DIAGNOSTIC ERROR:", err)
		http.Error(w, "Blueprint not found", http.StatusNotFound)
		return
	}

	tmpl, err := template.ParseFiles("templates/base.html", "templates/builder.html")
	if err != nil {
		log.Println("Template error:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	err = tmpl.ExecuteTemplate(w, "base", process)
	if err != nil {
		log.Println("TEMPLATE EXECUTION ERROR:", err)
		http.Error(w, "Template crash", http.StatusInternalServerError)
	}
}
