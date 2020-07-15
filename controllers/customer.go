package controllers

import (
	"net/http"

	"github.com/nurliman/Grasindo.API.go/config"
	"github.com/nurliman/Grasindo.API.go/models"

	"github.com/gin-gonic/gin"
)

// AddCustomer = adding costumer to database
func AddCustomer(context *gin.Context) {

	customer := models.Customer{Name: "hahahah"}
	config.DB.Create(&customer)
	context.JSON(http.StatusOK, gin.H{"data": customer})
}

// GetAllCustomers  Retrieve all customers
func GetAllCustomers(context *gin.Context) {
	var customers []models.Customer
	config.DB.Find(&customers)

	context.JSON(http.StatusOK, gin.H{"data": customers})
}

// GetACustomer find customer by id param
func GetACustomer(context *gin.Context) {
	var customer models.Customer
	config.DB.First(&customer, context.Param("id"))

	context.JSON(http.StatusOK, gin.H{"data": customer})
}

// DeleteACustomer delete customer by id param
func DeleteACustomer(context *gin.Context) {
	var customer models.Customer
	config.DB.First(&customer, context.Param("id"))
	config.DB.Delete(&customer)

	context.JSON(http.StatusOK, gin.H{"data": true})
}
