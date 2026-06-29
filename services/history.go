package services

import (
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
