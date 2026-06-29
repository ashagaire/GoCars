package services

import (
	"car-viewer/models"
	"car-viewer/utils"
	"sort"
)

const (
	categoryWeight     = 3
	manufacturerWeight = 2
	viewedPenalty      = 2
)

func BuildUserPreference(viewedCars []models.CarView) models.UserPreference {
	preference := models.UserPreference{
		Category:     make(map[string]int),
		Manufacturer: make(map[string]int),
		Drivetrain:   make(map[string]int),
		Transmission: make(map[string]int),
	}
	for _, car := range viewedCars {
		preference.Category[car.Category]++
		preference.Drivetrain[car.Specifications.Drivetrain]++
		preference.Manufacturer[car.Manufacturer]++
		preference.Transmission[car.Specifications.Transmission]++
	}
	return preference
}

func ScoreCar(car models.CarView, preference models.UserPreference) int {
	score := 0
	score += preference.Category[car.Category] * categoryWeight
	score += preference.Manufacturer[car.Manufacturer] * manufacturerWeight
	score += preference.Transmission[car.Specifications.Transmission]
	score += preference.Drivetrain[car.Specifications.Drivetrain]
	return score
}

func SortCarsByScore(scoredCars []models.ScoredCar) {
	sort.Slice(scoredCars, func(i, j int) bool {
		return scoredCars[i].Score > scoredCars[j].Score
	})
}

func TopCars(scoredCars []models.ScoredCar, limit int) []models.CarView {
	result := []models.CarView{}
	if len(scoredCars) < limit {
		limit = len(scoredCars)
	}
	for i := 0; i < limit; i++ {
		result = append(result, scoredCars[i].Car)
	}
	return result
}

func RecommendCars(currentCarID int, history []int, allCars []models.CarView, viewedCars []models.CarView) []models.CarView {
	if len(history) == 0 {
		history = []int{currentCarID}
	}

	preference := BuildUserPreference(viewedCars)
	scoredCars := make([]models.ScoredCar, 0, len(allCars))
	for _, car := range allCars {
		if car.ID == currentCarID {
			continue
		}

		scoredCar := models.ScoredCar{
			Car:   car,
			Score: ScoreCar(car, preference),
		}
		if utils.ContainsInt(history, car.ID) {
			scoredCar.Score /= viewedPenalty
		}

		scoredCars = append(scoredCars, scoredCar)
	}

	SortCarsByScore(scoredCars)
	return TopCars(scoredCars, 3)
}
