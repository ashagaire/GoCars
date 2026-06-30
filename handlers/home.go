package handlers

import (
	"car-viewer/models"
	"car-viewer/services"
	"car-viewer/utils"
	"html/template"
	"log"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	templates, err := template.New("").Funcs(template.FuncMap{
		"isSelected": isSelected,
		"isCompared": utils.ContainsInt,
	}).ParseGlob("templates/*.html")
	if err != nil {
		log.Printf("Error parsing templates: %v", err)

		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
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
	comparedIDs := utils.GetCompare(r)

	var filters models.CarFilters
	getQuery(r, &filters)
	filterCars := services.FilterCars(cars, filters)
	filterCarViews := services.BuildCarViews(filterCars, manufacturers, categories, comparedIDs)

	pageData := models.PageData{
		AllManufacture: allManufacturer,
		AllCategories:  allCategories,
		ActivePage:     "home",
		Cars:           filterCarViews,
		Manufacturers:  manufacturers,
		Categories:     categories,
		Filter:         filters,
		ComparedIDs:    comparedIDs,
	}

	err = templates.ExecuteTemplate(w, "home", pageData)

	if err != nil {
		log.Printf("ERROR: Executing template failed: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
