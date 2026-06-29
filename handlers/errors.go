package handlers

import (
	"log"
    "net/http"
	"html/template"

)

func InternalServerError(w http.ResponseWriter, r *http.Request, err error) {
	
	log.Fatalf("Error parsing templates: %v", err)
	
	w.WriteHeader(http.StatusInternalServerError)

	templates, err := template.ParseFiles("templates/errorPage.html")

	errTemplate := templates.ExecuteTemplate(w, "errorPage", nil)

	if errTemplate != nil {
		http.Error(w, errTemplate.Error(), http.StatusInternalServerError)
		return
	}
}
