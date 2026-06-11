package main

import (
	"fmt"
	"html/template"
	"net/http"
)

var tmplIndex = template.Must(template.ParseFiles("templates/index.html"))

func homeHandler(w http.ResponseWriter, r *http.Request) {
	tmplIndex.Execute(w, "Habg")
	cars, err := GetCars()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintln(w, "Car Viewer is running")
	for _, car := range cars {
		fmt.Fprintf(w, "%s (%d)\n", car.Name, car.Year)
	}

}
