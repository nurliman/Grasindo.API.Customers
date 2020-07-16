package models

import (
	"github.com/jinzhu/gorm"
)

// Customer containts customer data
type Customer struct {
	gorm.Model
	Name         string    `json:"name"`
	Addresses    []Address `json:"addresses"`
	Contacts     []Contact `json:"contacts"`
	OtherDetails string    `json:"otherDetails"`
}
