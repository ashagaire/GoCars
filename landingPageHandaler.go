package main

import (
	"html/template"
	"net/http"
)

type specificationsData struct {
    Engine string
    Horsepower int
    Transmission string
    Drivetrain string
}

type templateData struct {
	Id int
    Name string
	ManufacturerId string
    CategoryId string
    Year int
    Specifications specificationsData
    Image string
}

func landingPageHandaler(w http.ResponseWriter, req *http.Request) {
	tpl, err1 := template.ParseFiles("templates/landingPage.html")
	if  err1!= nil {
		http.Error(w, err1.Error(), http.StatusInternalServerError)
		return
	}
	data := templateData{}
	data.Id = 100
    data.Name = "Toyota Corolla"
	data.ManufacturerId = "1"
    data.CategoryId = "1"
    data.Year = 2023
    data.Specifications.Engine = "1.8L Inline-4"
    data.Specifications.Horsepower = 139
    data.Specifications.Transmission = "CVT"
    data.Specifications.Drivetrain = "Front-Wheel Drive"
    data.Image = "toyota_corolla.jpg"
	
	err := tpl.Execute(w,data)
	if err != nil {
		tpl.Execute(w, err.Error())
		return
	}
}