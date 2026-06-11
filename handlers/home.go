package handlers

import (
	"fmt"
	"html/template"
	"net/http"

	"../services/api"
)

var tmpl = template.Must(template.ParseFiles("templates/index.html"))

func homeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl.Execute(w, "Habg")
	cars, err := api.GetCars()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintln(w, "Car Viewer is running")
	for _, car := range cars {
		fmt.Fprintf(w, "%s (%d)\n", car.Name, car.Year)
	}
}
