package models

import "github.com/jinzhu/gorm"

type Location struct {
	gorm.Model
	Name      string  `json:"name" binding:"required"`
	Address   string  `json:"address" binding:"required"`
	Latitude  float64 `json:"latitude" binding:"required"`
	Longitude float64 `json:"longitude" binding:"required"`
	Category  string  `json:"category" binding:"required"`
	Distance  float64 `json:"distance,omitempty"`
}

type SearchParams struct {
	Latitude  float64 `json:"latitude" binding:"required"`
	Longitude float64 `json:"longitude" binding:"required"`
	Category  string  `json:"category" binding:"required"`
	RadiusKm  float64 `json:"radius_km" binding:"required"`
}

type OutputResponse struct {
    Locations []Location `json:"locations"`
}

type UserLocation struct {
	Latitude  float64 `json:"latitude" binding:"required"`
	Longitude float64 `json:"longitude" binding:"required"`
}

type TripCost struct {
	TotalCost float64 `json:"total_cost" binding:"required"`
	FuelCost  float64 `json:"fuel_cost" binding:"required"`
	TollCost  float64 `json:"toll_cost" binding:"required"`
}
