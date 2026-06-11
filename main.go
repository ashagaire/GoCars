package main

import (
	"fmt"
	"html/template"
	"net/http"
)

var tmpl = template.Must(template.ParseFiles("templates/index.html"),)

func main() {
	http.HandleFunc("/", homeHandler)
	fmt.Println("Go server running at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
