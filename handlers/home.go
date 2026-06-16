package handlers

import (
	"html/template"
	"net/http"

	"car-viewer/models"
	"car-viewer/services"
)

var Tmpl *template.Template

func HomeHandler(w http.ResponseWriter, r *http.Request) {

	cars, err := services.GetCars()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	manufacturers, err := services.GetManufacturers()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	categories, err := services.GetCategories()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	carViews := services.BuildCarViews(cars, manufacturers, categories)
	pageData := models.PageData{
		Cars: carViews,
		ActivePage: "home",
	}

	// homeTemplate.Execute(w, pageData)
	err = Tmpl.ExecuteTemplate(w, "home", pageData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}
