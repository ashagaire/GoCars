package services

import "car-viewer/models"

func BuildCarMap(cars []models.CarView) map[int]models.CarView {
	carMap := make(map[int]models.CarView)
	for _, car := range cars {
		carMap[car.ID] = car
	}
	return carMap
}

func BuildManufacturerMap(manufacturers []models.Manufacturer) map[int]models.Manufacturer {
	manufacturerMap := make(map[int]models.Manufacturer)
	for _, m := range manufacturers {
		manufacturerMap[m.ID] = m
	}
	return manufacturerMap
}

func BuildCategoryMap(categories []models.Category) map[int]string {
	categoryMap := make(map[int]string)
	for _, c := range categories {
		categoryMap[c.ID] = c.Name
	}
	return categoryMap
}
