package services

import (
	"car-viewer/models"
)

func BuildCarViews(cars []models.Car, manufacturers []models.Manufacturer, categories []models.Category) []models.CarView {
	manufacturerMap := BuildManufacturerMap(manufacturers)
	categoryMap := BuildCategoryMap(categories)

	var carViews []models.CarView
	for _, car := range cars {
		carView := models.CarView{
			ID:             car.ID,
			Name:           car.Name,
			Manufacturer:   manufacturerMap[car.ManufacturerID],
			Category:       categoryMap[car.CategoryID],
			Year:           car.Year,
			Specifications: car.Specifications,
			Image:          car.Image,
		}

		carViews = append(carViews, carView)
	}
	return carViews
}

func BuildCarDetailsView(carData models.Car, recommendedCars []models.CarView, manufacturers []models.Manufacturer, categories []models.Category) models.CarDetailView {
	manufacturerMap := BuildManufacturerMap(manufacturers)
	categoryMap := BuildCategoryMap(categories)

	carDetalsView := models.CarDetailView{
		ID:              carData.ID,
		Name:            carData.Name,
		Manufacturer:    manufacturerMap[carData.ManufacturerID],
		Category:        categoryMap[carData.CategoryID],
		Year:            carData.Year,
		ImageURL:        carData.Image,
		Specifications:  carData.Specifications,
		RecommendedCars: recommendedCars,
	}
	return carDetalsView
}

func BuildCarsCompareViews(cars []models.Car, manufacturers []models.Manufacturer, categories []models.Category) []models.CarView {
	manufacturerMap := BuildManufacturerMap(manufacturers)
	categoryMap := BuildCategoryMap(categories)

	var carViews []models.CarView
	for _, car := range cars {
		carView := models.CarView{
			ID:             car.ID,
			Name:           car.Name,
			Manufacturer:   manufacturerMap[car.ManufacturerID],
			Category:       categoryMap[car.CategoryID],
			Year:           car.Year,
			Specifications: car.Specifications,
			Image:          car.Image,
		}

		carViews = append(carViews, carView)
	}
	return carViews
}