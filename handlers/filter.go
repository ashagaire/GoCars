package handlers

import (
	"net/http"

	"car-viewer/models"
	"car-viewer/services"
)

func isSelected(ids []int, id int) bool {
	for _, selectedID := range ids {
		if selectedID == id {
			return true
		}
	}
	return false
}

func getQuery(r *http.Request, filters *models.CarFilters) {
	query := r.URL.Query().Get("q")
	yearFrom := r.URL.Query().Get("year_from")
	yearTo := r.URL.Query().Get("year_to")
	hpFrom := r.URL.Query().Get("hp_from")
	hpTo := r.URL.Query().Get("hp_to")
	transmission := r.URL.Query().Get("transmission")
	drivetrain := r.URL.Query().Get("drivetrain")
	manufacturerIDs := r.URL.Query()["manufacturer_id"]
	categoryID := r.URL.Query().Get("category_id")

	*filters = models.CarFilters{
		Query:           query,
		ManufacturerIDs: services.ParseIDs(manufacturerIDs),
		CategoryID:      services.ParseNumber(categoryID),
		YearFrom:        services.ParseNumber(yearFrom),
		YearTo:          services.ParseNumber(yearTo),
		HorsepowerFrom:  services.ParseNumber(hpFrom),
		HorsepowerTo:    services.ParseNumber(hpTo),
		Transmission:    transmission,
		Drivetrain:      drivetrain,
	}
}
