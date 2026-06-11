package services

import (
	"encoding/json"
	"net/http"

	"car-viewer/models"
)

func GetCars() ([]models.Car, error) {
	resp, err := http.Get("http://localhost:3000/api/models")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var cars []models.Car

	err = json.NewDecoder(resp.Body).Decode(&cars)
	if err != nil {
		return nil, err
	}
	return cars, nil
}

func GetManufacturers() ([]models.Manufacturer, error) {
	resp, err := http.Get("http://localhost:3000/api/manufacturers")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var manufacturers []models.Manufacturer

	err = json.NewDecoder(resp.Body).Decode(&manufacturers)
	if err != nil {
		return nil, err
	}

	return manufacturers, nil
}

func GetCategories() ([]models.Category, error) {
	resp, err := http.Get("http://localhost:3000/api/categories")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var categories []models.Category

	err = json.NewDecoder(resp.Body).Decode(&categories)
	if err != nil {
		return nil, err
	}

	return categories, nil
}
