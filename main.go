package main

import (
	"github.com/nurliman/Grasindo.API.go/config"
	"github.com/nurliman/Grasindo.API.go/models"
	"github.com/nurliman/Grasindo.API.go/routes"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/joho/godotenv/autoload"
)

var err error

func main() {

	config.DB, err = gorm.Open("postgres", config.DBConfigBuilder())

	if err != nil {
		panic("Failed to connect to database!")
	}

	defer config.DB.Close()

	config.DB.AutoMigrate(
		&models.Customer{},
		&models.Contact{},
		&models.Address{},
		&models.Coordinate{},
	)

	router := routes.SetupRouter()

	router.Run()
}
