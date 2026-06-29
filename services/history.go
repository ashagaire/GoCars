package services

import (
	"car-viewer/models"
	"car-viewer/utils"
	"net/http"
)

func UpdateHistory(w http.ResponseWriter, r *http.Request, carID int) []int {
	history := utils.GetHistory(r)
	history = utils.RemoveID(history, carID)
	history = append(history, carID)
	if len(history) > 5 {
		history = history[1:]
	}
	utils.SaveHistory(w, history)
	return history
}

func GetViewedCars(history []int, carMap map[int]models.CarView) []models.CarView {
	viewedCars := []models.CarView{}
	for _, id := range history {
		viewedCars = append(viewedCars, carMap[id])
	}
	return viewedCars
}
