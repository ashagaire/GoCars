package services

import (
	"car-viewer/models"
	"car-viewer/utils"
	"strconv"
	"strings"
)

func FilterCars(cars []models.Car, filters models.CarFilters) []models.Car {
	var filtered []models.Car
	for _, car := range cars {
		if !matchesQuery(car, filters.Query) {
			continue
		}
		if len(filters.ManufacturerIDs) != 0 && !utils.ContainsInt(filters.ManufacturerIDs, car.ManufacturerID) {
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
			!utils.ContainsIgnoreCase(car.Specifications.Transmission, filters.Transmission) {
			continue
		}
		if filters.Drivetrain != "" &&
			!utils.ContainsIgnoreCase(car.Specifications.Drivetrain, filters.Drivetrain) {
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
	return utils.ContainsIgnoreCase(car.Name, query) || utils.ContainsIgnoreCase(car.Specifications.Engine, query) || utils.ContainsIgnoreCase(strconv.Itoa(car.Specifications.Horsepower), query) || utils.ContainsIgnoreCase(car.Specifications.Transmission, query) || utils.ContainsIgnoreCase(car.Specifications.Drivetrain, query)
}
