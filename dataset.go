package main

type SpecificationsData struct {
	Engine       string
	Horsepower   int
	Transmission string
	Drivetrain   string
}

type TemplateData struct {
	Id             int
	Name           string
	ManufacturerId string
	CategoryId     string
	Year           int
	Specifications SpecificationsData
	Image          string
}

func dataset() []TemplateData {
	SpecificationCollection := SpecificationsData{

		Engine:       "1.8L Inline-4",
		Horsepower:   139,
		Transmission: "CVT",
		Drivetrain:   "Front-Wheel Drive",
	}
	DataCollection := []TemplateData{
		{
			Id:             1,
			Name:           "Toyota Corolla",
			ManufacturerId: "1",
			CategoryId:     "1",
			Year:           2023,
			Specifications: SpecificationCollection,
			Image:          "toyota_corolla.jpg",
		},
		{
			Id:             2,
			Name:           "Toyota Corolla",
			ManufacturerId: "1",
			CategoryId:     "1",
			Year:           2023,
			Specifications: SpecificationCollection,
			Image:          "toyota_corolla.jpg",
		},
		{
			Id:             3,
			Name:           "Toyota Corolla",
			ManufacturerId: "1",
			CategoryId:     "1",
			Year:           2023,
			Specifications: SpecificationCollection,
			Image:          "toyota_corolla.jpg",
		},
	}

	return DataCollection
}
