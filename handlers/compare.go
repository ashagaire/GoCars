package handlers

import (
	"car-viewer/models"
	"car-viewer/services"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

func ComparePageHandler(w http.ResponseWriter, r *http.Request) {

	templates, err := template.New("").Funcs(template.FuncMap{
		"isSelected": isSelected,
	}).ParseGlob("templates/*.html")
	if err != nil {
		log.Printf("Error parsing templates: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	var req CompareRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

	carsData, err := services.GetCompareCars(req)
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

	carViews := services.BuildCarsCompareViews(carsData, manufacturers, categories)
	comparePageData := models.ComparePageData{
		CompareCars:     carViews,
		ActivePage:     "compare",
	}

	err = templates.ExecuteTemplate(w, "Compare", comparePageData)
	if err != nil {
		log.Printf("ERROR: Executing template failed: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}
