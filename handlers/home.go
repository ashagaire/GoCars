package handlers

import (
	"html/template"
	"net/http"
	"log"
	"car-viewer/models"
	"car-viewer/services"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {

	templates, err := template.ParseGlob("templates/*.html")
	if err != nil {
		log.Fatalf("Error parsing templates: %v", err)
	}

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

	allManufacturer := len(manufacturers)
	allCategories:= len(categories)
	
	carViews := services.BuildCarViews(cars, manufacturers, categories)
	pageData := models.PageData{
		Cars: carViews,
		AllManufacture: allManufacturer,
		AllCategories: allCategories,
		ActivePage: "home",
	}

	err = templates.ExecuteTemplate(w, "home", pageData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}
