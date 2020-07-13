package models

import (
	"github.com/jinzhu/gorm"

	// register the driver
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// DB containts information for current DB connection
var DB *gorm.DB

// ConnectDataBase = connecting to database
func ConnectDataBase() {
	db, err := gorm.Open("postgres", "host=localhost port=5432 user=mypgadmin dbname=db0 password=vwxyz123")

	if err != nil {
		panic("Failed to connect to database!")
	}

	defer db.Close()

	db.AutoMigrate(&Customer{})

	DB = db
}
