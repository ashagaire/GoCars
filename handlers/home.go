package handlers

import (
	"fmt"
	"html/template"
	"net/http"

	"car-viewer/services"
)

var tmpl = template.Must(template.ParseFiles("templates/index.html"))

func HomeHandler(w http.ResponseWriter, r *http.Request) {

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

	carViews := services.BuildCarViews(cars, manufacturers, categories)

	tmpl.Execute(w, carViews)

	fmt.Fprintln(w, "Car Viewer is running")
	for _, car := range cars {
		fmt.Fprintf(w, "%s (%d)\n", car.Name, car.Year)
	}
}
