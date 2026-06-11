package services

import (
	"encoding/json"
	"net/http"

	c "../models/car"
)

func GetCars() ([]c.Car, error) {
	resp, err := http.Get("http://localhost:3000/api/models")
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var cars []c.Car

	err = json.NewDecoder(resp.Body).Decode(&cars)

	if err != nil {
		return nil, err
	}

	return cars, nil
}
