package models

import (
	"github.com/jinzhu/gorm"
)

// Coordinate containts Latitude and Longitude
type Coordinate struct {
	gorm.Model
	Latitude  float64 `json:"lat" binding:"required"`
	Longitude float64 `json:"lng" binding:"required"`
	AddressID uint
}
