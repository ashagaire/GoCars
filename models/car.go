package models

type Car struct {
	ID             int            `json:"id"`
	Name           string         `json:"name"`
	ManufacturerID int            `json:"manufacturerId"`
	CategoryID     int            `json:"categoryId"`
	Year           int            `json:"year"`
	Specifications Specifications `json:"specifications"`
	Image          string         `json:"image"`
}

type Specifications struct {
	Engine       string `json:"engine"`
	Horsepower   int    `json:"horsepower"`
	Transmission string `json:"transmission"`
	Drivetrain   string `json:"drivetrain"`
}

type Manufacturer struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Country      string `json:"country"`
	FoundingYear int    `json:"foundingYear"`
}

type Category struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type CarView struct {
	ID             int
	Name           string
	Manufacturer   string
	Category       string
	Year           int
	Specifications Specifications
	Image          string
}

type PageData struct {
	Cars           []CarView
	ActivePage     string
	AllManufacture int
	AllCategories  int
	// 	SearchQuery string

	// 	SelectedCategory string

	// Error string
	Manufacturers []Manufacturer
	Categories    []Category
	Filter        CarFilters
}

type CarFilters struct {
	Query           string
	ManufacturerIDs []int
	CategoryID      int
	YearFrom        int
	YearTo          int
	HorsepowerFrom  int
	HorsepowerTo    int
	Transmission    string
	Drivetrain      string
}

type CarDetailView struct {
	ID              int
	Name            string
	Manufacturer    string
	Category        string
	Year            int
	ImageURL        string
	Specifications  Specifications
	RecommendedCars []CarView
}

//type CompareView struct {}

//type RecommendationView struct {}

type ScoredCar struct {
	Car   CarView
	Score int
}
