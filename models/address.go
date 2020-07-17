package models

import (
	"github.com/jinzhu/gorm"
)

// Address containts address(es) of customer
type Address struct {
	gorm.Model
	Name         string      `json:"name" binding:"required"`
	AddressLine1 string      `json:"addressLine1"`
	AddressLine2 string      `json:"addressLine2"`
	SubDistrict  string      `json:"subDistrict"`
	Village      string      `json:"village"`
	City         string      `json:"city"`
	Province     string      `json:"province"`
	Country      string      `json:"country"`
	ZipCode      int         `json:"zipCode"`
	Coordinate   *Coordinate `json:"coordinate" binding:"required,dive"`
	OtherDetails string      `json:"otherDetails"`
	CustomerID   uint
}
