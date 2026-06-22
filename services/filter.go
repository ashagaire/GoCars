package services

import (
	"car-viewer/models"
	"strconv"
	"strings"
)

func FilterCars(cars []models.Car, filters models.CarFilters) []models.Car {
	var filtered []models.Car
	for _, car := range cars {
		if len(filters.ManufacturerIDs) != 0 && !containsInt(filters.ManufacturerIDs, car.ManufacturerID) {
			continue
		}
		if filters.CategoryIDs != 0 && filters.CategoryIDs != car.CategoryID {
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

func ParseIDs(strIDs []string) []int {
	var ids []int
	for _, strID := range strIDs {
		id, err := strconv.Atoi(strID)
		if err == nil {
			ids = append(ids, id)
		}
	}
	return ids
}

func ParseNumber(strNumber string) int {
	if strNumber == "" {
		return 0
	}
	number, err := strconv.Atoi(strNumber)
	if err != nil {
		return 0
	}
	return number
}
