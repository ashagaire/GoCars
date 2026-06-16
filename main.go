package main

import (
	"car-viewer/handlers"
	"fmt"
	"net/http"
	"html/template"
	"log"
)

func main() {
	// 1. Parse all UI templates once at startup
	var err error
	handlers.Tmpl, err = template.ParseGlob("templates/*.html")
	if err != nil {
		log.Fatalf("Error parsing templates: %v", err)
	}

	fileServer := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static", fileServer))

	http.HandleFunc("/", handlers.HomeHandler)
	fmt.Println("Go server running at http://localhost:8080")
	err1 := http.ListenAndServe(":8080", nil)
	log.Fatalln(err1)
}

