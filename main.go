package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fileServer := http.FileServer(http.Dir("static"))
	http.Handle("/ststiic/", http.StripPrefix("/static", fileServer))

	http.HandleFunc("/", landingPageHandaler)
	fmt.Println("Server is running  at http://localhost:8080 ...")

	err := http.ListenAndServe(":8080", nil)
	log.Fatalln(err)
}
