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

func CarDetailsPageHandler(w http.ResponseWriter, r *http.Request) {

	templates, err := template.New("").Funcs(template.FuncMap{
		"isSelected": isSelected,
	}).ParseGlob("templates/*.html")
	if err != nil {
		log.Printf("Error parsing templates: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}
	carIdStr := r.URL.Query().Get("id")

	comparedMap := make(map[int]bool)
	if cookie, err := r.Cookie("compare_cars"); err == nil && cookie.Value != "" {
		comparedIDs := strings.Split(cookie.Value, ",")
		for _, id := range comparedIDs {
			if idInt, err := strconv.Atoi(id); err == nil {
            	comparedMap[idInt] = true
       		}
		}
	}

	carData, err := services.GetCarbyID(carIdStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	carID, err := strconv.Atoi(carIdStr)
	if err != nil {
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

	allCars := services.BuildCarViews(cars, manufacturers, categories, comparedMap)
	history := services.UpdateHistory(w, r, carID)
	carMap := services.BuildCarMap(allCars)
	viewedCars := services.GetViewedCars(history, carMap)
	recommendedCars := services.RecommendCars(carID, history, allCars, viewedCars)
	manufacturerMap := services.BuildManufacturerMap(manufacturers)
	currentMfg := manufacturerMap[carData.ManufacturerID]

	carViews := services.BuildCarDetailsView(carData, recommendedCars, viewedCars, manufacturers, categories)
	carDetailPageData := models.CarDetailView{
		ID:              carViews.ID,
		Name:            carViews.Name,
		ManufacturerName:    currentMfg.Name,
		ManufacturerCountry: currentMfg.Country,
		ManufacturerYear:	currentMfg.FoundingYear,
		Category:        carViews.Category,
		Year:            carViews.Year,
		ImageURL:        carViews.ImageURL,
		Specifications:  carViews.Specifications,
		RecommendedCars: carViews.RecommendedCars,
		RecentCars:      carViews.RecentCars,
	}

	err = templates.ExecuteTemplate(w, "CarDetails", carDetailPageData)
	if err != nil {
		log.Printf("ERROR: Executing template failed: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}
