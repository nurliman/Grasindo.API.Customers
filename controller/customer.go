package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nurliman/Grasindo.API.go/models"
)

// AddCustomer = adding costumer to database
func AddCustomer(context *gin.Context) {

	customer := models.Customer{Name: "hahahah"}
	models.DB.Create(&customer)
	context.JSON(http.StatusOK, gin.H{"data": customer})
}
