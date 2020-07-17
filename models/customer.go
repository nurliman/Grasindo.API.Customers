package models

import (
	"github.com/jinzhu/gorm"
)

// Customer containts customer data
type Customer struct {
	gorm.Model
	Name         string    `json:"name" binding:"required"`
	Addresses    []Address `json:"addresses" binding:"required,min=1,dive"`
	Contacts     []Contact `json:"contacts" binding:"required,min=1,dive"`
	OtherDetails string    `json:"otherDetails"`
}

// EditCustomerInput = data model when Add/Edit Customer
type EditCustomerInput struct {
	Name         string `json:"name" binding:"required"`
	OtherDetails string `json:"otherDetails"`
}
