package models

import (
	"github.com/jinzhu/gorm"

	// register the driver
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// DB containts information for current DB connection
var DB *gorm.DB

// ConnectDataBase = connecting to database
func ConnectDataBase() {
	db, err := gorm.Open("sqlite3", "test.db")

	if err != nil {
		panic("Failed to connect to database!")
	}

	defer db.Close()

	db.AutoMigrate(&Customer{})

	DB = db
}
