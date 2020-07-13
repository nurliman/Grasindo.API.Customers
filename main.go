package main

import (
	"github.com/nurlimandiara/Grasindo.API.go/models"

	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	router.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{"data": "Hello World!"})
	})

	models.ConnectDataBase()

	router.Run()
}
