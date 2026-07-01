package main

import (
	"car-viewer/handlers"
	"fmt"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	fileServer := http.FileServer(http.Dir("static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fileServer))

	imageServer := http.FileServer(http.Dir("./api/img"))
	mux.Handle("/api-car-images/", http.StripPrefix("/api-car-images/", imageServer))

	assetsServer := http.FileServer(http.Dir("assets"))
	mux.Handle("/assets/", http.StripPrefix("/assets/", assetsServer))

	mux.HandleFunc("/", handlers.HomeHandler)
	mux.HandleFunc("/car", handlers.CarDetailsPageHandler)
	mux.HandleFunc("/manufacturers", handlers.ManufacturerPageHandler)
	mux.HandleFunc("/compare", handlers.ComparePageHandler)

	fmt.Println("Go server running at http://localhost:8080")
	err1 := http.ListenAndServe(":8080", mux)
	log.Fatalln(err1)
}
