package models

import (
	"github.com/jinzhu/gorm"
)

// Coordinate containts Latitude and Longitude
type Coordinate struct {
	gorm.Model
	Latitude  float64 `json:"lat"`
	Longitude float64 `json:"lng"`
	AddressID uint
}
