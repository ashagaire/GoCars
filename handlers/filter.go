package handlers

import (
	"car-viewer/models"
	"car-viewer/utils"
	"net/http"
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
		ManufacturerIDs: utils.ParseIDs(manufacturerIDs),
		CategoryID:      utils.ParseNumber(categoryID),
		YearFrom:        utils.ParseNumber(yearFrom),
		YearTo:          utils.ParseNumber(yearTo),
		HorsepowerFrom:  utils.ParseNumber(hpFrom),
		HorsepowerTo:    utils.ParseNumber(hpTo),
		Transmission:    transmission,
		Drivetrain:      drivetrain,
	}
}
