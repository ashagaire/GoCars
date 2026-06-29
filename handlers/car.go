package handlers

import (
	"car-viewer/models"
	"car-viewer/services"
	"html/template"
	"log"
	"net/http"
)

func CarDetailsPageHandler(w http.ResponseWriter, r *http.Request) {

	templates, err := template.New("").Funcs(template.FuncMap{
		"isSelected": isSelected,
	}).ParseGlob("templates/*.html")
	if err != nil {
		log.Fatalf("Error parsing templates: %v", err)
	}
	carId := r.URL.Query().Get("id")

	carData, err := services.GetCarbyID(carId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	recommendedCars, err := services.GetRecommendedCars()
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

	carViews := services.BuildCarDetailsViews(carData, recommendedCars, manufacturers, categories)
	carDetailPageData := models.CarDetailView{
		ID:             carViews.ID,
		Name:           carViews.Name,
		Manufacturer:   carViews.Manufacturer,
		Category:       carViews.Category,
		Year:             carViews.Year,
		ImageURL:         carViews.ImageURL,
		Specifications: carViews.Specifications,
		RecommendedCars: carViews.RecommendedCars,
	}

	err = templates.ExecuteTemplate(w, "CarDetails", carDetailPageData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}