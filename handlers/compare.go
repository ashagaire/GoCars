package handlers

import (
	"car-viewer/models"
	"car-viewer/services"
	"car-viewer/utils"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func CompareHandler(w http.ResponseWriter, r *http.Request) {
	templates, err := template.New("").Funcs(template.FuncMap{
		"isSelected": isSelected,
		"isCompared": utils.ContainsInt,
	}).ParseGlob("templates/*.html")
	if err != nil {
		log.Printf("Error parsing templates: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	cars, _ := services.GetCars()
	manufacturers, _ := services.GetManufacturers()
	categories, _ := services.GetCategories()
	compareIDs := utils.GetCompare(r)
	compareCars := services.GetCompareCars(cars, compareIDs)
	compareViews := services.BuildCarViews(
		compareCars,
		manufacturers,
		categories,
		compareIDs,
	)
	templates.ExecuteTemplate(w, "comparison", models.PageData{
		Cars:       compareViews,
		ActivePage: "compare",
	})
}

func AddCompareHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	id := utils.ParseNumber(r.FormValue("id"))
	compare := services.UpdateCompare(w, r, id)
	returnURL := r.Referer()
	if returnURL == "" {
		returnURL = "/"
	}
	fmt.Println(compare)

	http.Redirect(w, r, returnURL, http.StatusSeeOther)
}

func RemoveCompareHandler(w http.ResponseWriter, r *http.Request) {
	id := utils.ParseNumber(r.FormValue("id"))
	services.RemoveCompare(w, r, id)
	returnURL := r.Referer()
	if returnURL == "" {
		returnURL = "/"
	}
	http.Redirect(w, r, returnURL, http.StatusSeeOther)
}
