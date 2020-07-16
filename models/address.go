package models

import (
	"github.com/jinzhu/gorm"
)

// Address containts address(es) of customer
type Address struct {
	gorm.Model
	Name         string      `json:"name"`
	AddressLine1 string      `json:"addressLine1"`
	AddressLine2 string      `json:"addressLine2"`
	SubDistrict  string      `json:"subDistrict"`
	Village      string      `json:"village"`
	City         string      `json:"city"`
	Province     string      `json:"province"`
	Country      string      `json:"country"`
	ZipCode      int         `json:"zipCode"`
	Coordinate   *Coordinate `json:"coordinate"`
	OtherDetails string      `json:"otherDetails"`
	CustomerID   uint
}
