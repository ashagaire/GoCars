package services

import "car-viewer/models"

func RecommendCars(currentCar models.Car, history []int) {
	if len(history) == 0 {
		history = append(history, currentCar.ID)
	}
}
