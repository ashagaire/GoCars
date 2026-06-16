package main

import (
	"car-viewer/handlers"
	"fmt"
	"net/http"
	"log"
)

func main() {
	fileServer := http.FileServer(http.Dir("static"))
	http.Handle("/ststiic/", http.StripPrefix("/static", fileServer))

	http.HandleFunc("/", handlers.HomeHandler)
	fmt.Println("Go server running at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
	log.Fatalln(err)
}

