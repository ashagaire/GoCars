package main

import (
	"encoding/json"
	"net/http"
)

type Car struct {
	ID             int            `json:"id"`
	Name           string         `json:"name"`
	ManufacturerID int           `json:"manufacturerId"`
	CategoryID     int           `json:"categoryId"`
	Year           int           `json:"year"`
	Specifications Specifications `json:"specifications"`
	Image          string         `json:"image"`
}

type Specifications struct {
	Engine       string `json:"engine"`
	Horsepower   int    `json:"horsepower"`
	Transmission string `json:"transmission"`
	Drivetrain   string `json:"drivetrain"`
}

func GetCars() ([]Car, error) {
	resp, err := http.Get("http://localhost:3000/api/models")
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var cars []Car

	err = json.NewDecoder(resp.Body).Decode(&cars)
	
	if err != nil {
		return nil, err
	}

	return cars, nil
}