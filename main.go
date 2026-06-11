package main

import (
	"fmt"
	"net/http"
	h "/handlers/home"
)

func main() {
	http.HandleFunc("/", h.homeHandler)
	fmt.Println("Go server running at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
