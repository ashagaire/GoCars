package models

type UserPreference struct {
    Category map[string]int
    Manufacturer map[string]int
    Transmission map[string]int
    Drivetrain map[string]int
}