package services

import (
	"car-viewer/models"
	"encoding/json"
	"strconv"
	"net/http"
)

func GetCompareCars(ids []int) ([]models.Car, error) {
	var listCars []models.Car
	
	for _, id := range ids {
		idStr := strconv.Itoa(id)
		resp, err := http.Get("http://localhost:3000/api/models/" + idStr)
		if err != nil {
			return listCars, err
		}
		defer resp.Body.Close()

		var car models.Car
		err = json.NewDecoder(resp.Body).Decode(&car)
		if err != nil {
			return listCars, err
		}
		listCars = append(listCars, car )
	}
	return listCars, nil
}