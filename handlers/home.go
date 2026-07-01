package handlers

import (
	"car-viewer/models"
	"car-viewer/services"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	templates, err := template.New("").Funcs(template.FuncMap{
		"isSelected": isSelected,
	}).ParseGlob("templates/*.html")
	if err != nil {
		log.Printf("Error parsing templates: %v", err)

		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	comparedMap := make(map[int]bool)
	if cookie, err := r.Cookie("compare_cars"); err == nil && cookie.Value != "" {
		comparedIDs := strings.Split(cookie.Value, ",")
		for _, id := range comparedIDs {
			if idInt, err := strconv.Atoi(id); err == nil {
				comparedMap[idInt] = true
			}
		}
	}

	manufacturers, categories, cars, err := services.FetchAllData()
	if err != nil {
		ServerError(w, "Fetching data failed", err)
		return
	}

	allManufacturer := len(manufacturers)
	allCategories := len(categories)

	var filters models.CarFilters
	getQuery(r, &filters)
	filterCars := services.FilterCars(cars, filters)

	filterCarViews := services.BuildCarViews(filterCars, manufacturers, categories, comparedMap)

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
		log.Printf("ERROR: Executing template failed: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func ServerError(w http.ResponseWriter, logMsg string, err error) {
	log.Printf("Error: %s : %v", logMsg, err)
	http.Error(w, "Oops, something went wrong. Please try again later.", http.StatusInternalServerError)
}
