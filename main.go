package main

import (
	"car-viewer/handlers"
	"fmt"
	"log"
	"net/http"
)

func main() {
	fileServer := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fileServer))

	imageServer := http.FileServer(http.Dir("./api/img"))
	http.Handle("/api-car-images/", http.StripPrefix("/api-car-images/", imageServer))

	assetsServer := http.FileServer(http.Dir("assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", assetsServer))

	http.HandleFunc("/", handlers.HomeHandler)
	http.HandleFunc("/car", handlers.CarDetailsPageHandler)
	http.HandleFunc("/manufacturers", handlers.ManufacturerPageHandler)
	fmt.Println("Go server running at http://localhost:8080")
	err1 := http.ListenAndServe(":8080", nil)
	log.Fatalln(err1)
}

