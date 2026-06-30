package handlers

import (
	"car-viewer/models"
	"car-viewer/services"
	"html/template"
	"log"
	"strconv"
	"strings"
	"net/http"
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

	var ids []int

	cookie, err := r.Cookie("compare_cars")
	if err == nil && cookie.Value != "" {
		idStrings := strings.Split(cookie.Value, ",")
		for _, id := range idStrings {
			if idInt, err := strconv.Atoi(id); err == nil {
				ids = append(ids, idInt)
			}
		}
	}

	if len(ids) == 0 {
		 http.Redirect(w, r, "/", http.StatusSeeOther)
        return
	}


	carsData, err := services.GetCompareCars(ids)
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
