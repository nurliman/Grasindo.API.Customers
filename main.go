package main

import (
	"github.com/nurliman/Grasindo.API.go/models"
	"github.com/nurliman/Grasindo.API.go/routes"
)

func main() {

	models.ConnectDataBase()

	router := routes.SetupRouter()

	router.Run()
}
