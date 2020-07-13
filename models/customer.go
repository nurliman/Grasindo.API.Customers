package models

import (
	"github.com/jinzhu/gorm"
)

// Customer containts customer data
type Customer struct {
	gorm.Model
	Name         string    `json:"name"`
	Adress       []Address `json:"addresses"`
	Contact      []Contact `json:"contacts"`
	OtherDetails string    `json:"otherDetails"`
}
