package handlers

import (
	"car-viewer/models"
	"car-viewer/services"
	"html/template"
	"log"
	"net/http"
)

func ManufacturerPageHandler(w http.ResponseWriter, r *http.Request) {
	// templates, err := template.ParseGlob("templates/*.html")

	templates, err := template.New("").Funcs(template.FuncMap{
		"isSelected": isSelected,
	}).ParseGlob("templates/*.html")
	if err != nil {
		log.Printf("Error parsing templates: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}
	

	allManufacturers, err := services.GetManufacturers()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	totalManufacturer := len(allManufacturers)
	
	manufacturerData := models.ManufacturerData{
		AllManufacture: allManufacturers,
		TotalManufacturer: totalManufacturer,
		ActivePage:     "manufacturer",
	}

	err = templates.ExecuteTemplate(w, "Manufacturar", manufacturerData)
	if err != nil {
		log.Printf("ERROR: Executing template failed: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}
