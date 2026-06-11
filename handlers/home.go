package handlers

import (
	"html/template"
	"net/http"

	"car-viewer/models"
	"car-viewer/services"
)

var tmpl = template.Must(template.ParseFiles("templates/index.html"))

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
	}

	tmpl.Execute(w, pageData)

}
