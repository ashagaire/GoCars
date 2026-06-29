package handlers

import (
	"car-viewer/models"
	"car-viewer/services"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

func CarDetailsPageHandler(w http.ResponseWriter, r *http.Request) {
	// templates, err := template.ParseGlob("templates/*.html")

	templates, err := template.New("").Funcs(template.FuncMap{
		"isSelected": isSelected,
	}).ParseGlob("templates/*.html")
	if err != nil {
		log.Printf("Error parsing templates: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}
	carIdStr := r.URL.Query().Get("id")

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

	allCars := services.BuildCarViews(cars, manufacturers, categories)
	history := services.UpdateHistory(w, r, carID)

	recommendedCars := services.RecommendCars(carID, history, allCars)

	carViews := services.BuildCarDetailsView(carData, recommendedCars, manufacturers, categories)
	carDetailPageData := models.CarDetailView{
		ID:              carViews.ID,
		Name:            carViews.Name,
		Manufacturer:    carViews.Manufacturer,
		Category:        carViews.Category,
		Year:            carViews.Year,
		ImageURL:        carViews.ImageURL,
		Specifications:  carViews.Specifications,
		RecommendedCars: carViews.RecommendedCars,
	}

	err = templates.ExecuteTemplate(w, "CarDetails", carDetailPageData)
	if err != nil {
		log.Printf("ERROR: Executing template failed: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}
