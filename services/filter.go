package services

import (
	"car-viewer/models"
	"strings"
)

func FilterCars(cars []models.Car, filters models.CarFilters) []models.Car {
	var filtered []models.Car
	for _, car := range cars {
		if len(filters.ManufacturerIDs) != 0 && !containsInt(filters.ManufacturerIDs, car.ManufacturerID) {
			continue
		}
		if len(filters.CategoryIDs) != 0 && !containsInt(filters.CategoryIDs, car.CategoryID) {
			continue
		}
		if filters.YearFrom != 0 && car.Year < filters.YearFrom {
			continue
		}
		if filters.YearTo != 0 && car.Year > filters.YearTo {
			continue
		}
		if !matchesQuery(car, filters.Query) {
			continue
		}
		filtered = append(filtered, car)
	}
	return filtered
}

func containsInt(ints []int, n int) bool {
	for _, i := range ints {
		if i == n {
			return true
		}
	}
	return false
}

func containsIgnoreCase(str, substr string) bool {
	return strings.Contains(strings.ToLower(str), strings.ToLower(substr))
}

func matchesQuery(car models.Car, query string) bool {
	query = strings.TrimSpace(query)
	if query == "" {
		return true
	}
	return containsIgnoreCase(car.Name, query) || containsIgnoreCase(car.Specifications.Engine, query) || containsIgnoreCase(car.Specifications.Transmission, query) || containsIgnoreCase(car.Specifications.Drivetrain, query)
}
