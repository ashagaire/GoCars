package services

import (
	"car-viewer/models"
	"encoding/json"
	"net/http"
)

func GetCompareCars(ids []string) ([]models.Car, error) {
	listCars := []models.Car
	if ids == nil {
		return nil, err
	}
	for id := range ids {
		resp, err := http.Get("http://localhost:3000/api/models/" + id)
		if err != nil {
			return models.Car{}, err
		}
		defer resp.Body.Close()

		var car models.Car
		err = json.NewDecoder(resp.Body).Decode(&car)
		if err != nil {
			return models.Car{}, err
		}
		listCars = append(listCars, car )
	}
	return listCars, nil
}