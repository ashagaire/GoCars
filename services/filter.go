package services

import (
	"car-viewer/models"
	"strconv"
	"strings"
)

func FilterCars(cars []models.Car, filters models.CarFilters) []models.Car {
	var filtered []models.Car
	for _, car := range cars {
		if !matchesQuery(car, filters.Query) {
			continue
		}
		if len(filters.ManufacturerIDs) != 0 && !utils.containsInt(filters.ManufacturerIDs, car.ManufacturerID) {
			continue
		}
		if filters.CategoryID != 0 && filters.CategoryID != car.CategoryID {
			continue
		}
		if filters.YearFrom != 0 && car.Year < filters.YearFrom {
			continue
		}
		if filters.YearTo != 0 && car.Year > filters.YearTo {
			continue
		}
		if filters.HorsepowerFrom != 0 &&
			car.Specifications.Horsepower < filters.HorsepowerFrom {
			continue
		}
		if filters.HorsepowerTo != 0 &&
			car.Specifications.Horsepower > filters.HorsepowerTo {
			continue
		}
		if filters.Transmission != "" &&
			!utils.containsIgnoreCase(car.Specifications.Transmission, filters.Transmission) {
			continue
		}
		if filters.Drivetrain != "" &&
			!utils.containsIgnoreCase(car.Specifications.Drivetrain, filters.Drivetrain) {
			continue
		}
		filtered = append(filtered, car)
	}
	return filtered
}

func matchesQuery(car models.Car, query string) bool {
	query = strings.TrimSpace(query)
	if query == "" {
		return true
	}
	return utils.containsIgnoreCase(car.Name, query) || utils.containsIgnoreCase(car.Specifications.Engine, query) || utils.containsIgnoreCase(strconv.Itoa(car.Specifications.Horsepower), query) || utils.containsIgnoreCase(car.Specifications.Transmission, query) || utils.containsIgnoreCase(car.Specifications.Drivetrain, query)
}
