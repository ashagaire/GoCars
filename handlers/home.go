package handlers

import (
	"car-viewer/models"
	"car-viewer/services"
	"html/template"
	"log"
	"net/http"
)

var homeTemplate = template.Must(
	template.New("home.html").Funcs(template.FuncMap{
		"isSelected": isSelected,
	}).ParseFiles("templates/home.html"))

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
	allCategories := len(categories)

	query := r.URL.Query().Get("q")
	yearFrom := r.URL.Query().Get("year_from")
	yearTo := r.URL.Query().Get("year_to")
	manufacturerIDs := r.URL.Query()["manufacturer_id"]
	categoryIDs := r.URL.Query()["category_id"]

	filters := models.CarFilters{
		Query:           query,
		ManufacturerIDs: services.ParseIDs(manufacturerIDs),
		CategoryIDs:     services.ParseIDs(categoryIDs),
		YearFrom:        services.ParseNumber(yearFrom),
		YearTo:          services.ParseNumber(yearTo),
	}

	filterCars := services.FilterCars(cars, filters)
	filterCarViews := services.BuildCarViews(filterCars, manufacturers, categories)

	pageData := models.PageData{
		AllManufacture: allManufacturer,
		AllCategories:  allCategories,
		ActivePage:     "home",
		Cars:           filterCarViews,
		Manufacturers:  manufacturers,
		Categories:     categories,
		Filter:         filters,
	}

	err = templates.ExecuteTemplate(w, "home", pageData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
