package models

import (
	"github.com/jinzhu/gorm"
)

// Contact containts contact(s) of customer
type Contact struct {
	gorm.Model
	Name         string `json:"name" binding:"required"`
	Title        string `json:"title"`
	Phone        string `json:"phone" binding:"required"`
	Phone2       string `json:"phone2"`
	Phone3       string `json:"phone3"`
	Email        string `json:"email"`
	OtherDetails string `json:"otherDetails"`
	CustomerID   uint
}

type ContactInput struct {
	Name         string `json:"name" binding:"required"`
	Title        string `json:"title"`
	Phone        string `json:"phone" binding:"required"`
	Phone2       string `json:"phone2"`
	Phone3       string `json:"phone3"`
	Email        string `json:"email"`
	OtherDetails string `json:"otherDetails"`
}
