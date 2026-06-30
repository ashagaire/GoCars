package services

import (
	"car-viewer/models"
	"car-viewer/utils"
	"net/http"
)

func UpdateCompare(w http.ResponseWriter, r *http.Request, carID int) []int {
	compare := utils.GetCompare(r)
	compare = utils.RemoveID(compare, carID)
	compare = append(compare, carID)
	if len(compare) > 4 {
		compare = compare[1:]
	}
	utils.SaveCompare(w, compare)
	return compare
}

func RemoveCompare(w http.ResponseWriter, r *http.Request, carID int) {
	compare := utils.GetCompare(r)
	compare = utils.RemoveID(compare, carID)
	utils.SaveCompare(w, compare)
}

func GetCompareCars(cars []models.Car, ids []int) []models.Car {

	var result []models.Car
	for _, id := range ids {
		for _, car := range cars {
			if car.ID == id {
				result = append(result, car)
				break
			}
		}
	}
	return result
}
