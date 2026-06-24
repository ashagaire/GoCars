package services

import (
	"car-viewer/models"
)

func BuildCarViews(cars []models.Car, manufacturers []models.Manufacturer, categories []models.Category) []models.CarView {
	manufacturerMap := make(map[int]string)
	for _, m := range manufacturers {
		manufacturerMap[m.ID] = m.Name
	}

	categoryMap := make(map[int]string)
	for _, c := range categories {
		categoryMap[c.ID] = c.Name
	}

	var carViews []models.CarView
	for _, car := range cars {
		carView := models.CarView{
			ID:           car.ID,
			Name:         car.Name,
			Manufacturer: manufacturerMap[car.ManufacturerID],
			Category:     categoryMap[car.CategoryID],
			Year:         car.Year,
			Specifications: car.Specifications,
			Image:        car.Image,
		}

		carViews = append(carViews, carView)
	}
	return carViews
}
