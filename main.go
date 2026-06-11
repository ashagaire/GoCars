package main

import (
	"car-viewer/handlers"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", handlers.HomeHandler)
	fmt.Println("Go server running at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
