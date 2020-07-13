package models

import (
	"github.com/jinzhu/gorm"
)

// Coordinate containts Latitude and Longitude
type Coordinate struct {
	Latitude  float64 `json:"lat"`
	Longitude float64 `json:"lng"`
}

// Address containts address(es) of customer
type Address struct {
	gorm.Model
	Name         string     `json:"name"`
	AdressLine1  string     `json:"addressLine1"`
	AdressLine2  string     `json:"addressLine2"`
	SubDistrict  string     `json:"subDistrict"`
	Village      string     `json:"village"`
	City         string     `json:"city"`
	Province     string     `json:"province"`
	Country      string     `json:"country"`
	ZipCode      int        `json:"zipCode"`
	Coordinate   Coordinate `json:"coordinate"`
	OtherDetails string     `json:"otherDetails"`
}
