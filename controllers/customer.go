package controllers

import (
	"fmt"
	"net/http"

	"github.com/nurliman/Grasindo.API.go/config"
	"github.com/nurliman/Grasindo.API.go/models"

	"github.com/gin-gonic/gin"
)

// AddCustomer = adding costumer to database
func AddCustomer(context *gin.Context) {
	var input models.Customer
	if err := context.BindJSON(&input); err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config.DB.Create(&input)
	context.JSON(http.StatusOK, gin.H{"data": input})
}

// EditACustomer update edit customer information, but not updating sub models like Address, Customer
func EditACustomer(context *gin.Context) {
	var customer models.Customer
	var id = context.Param("id")

	if err := config.DB.
		Preload("Addresses").
		Preload("Addresses.Coordinate").
		Preload("Contacts").
		Where("id = ?", id).
		First(&customer).
		Error; err != nil {
		context.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"message": fmt.Sprintf("Cannot find customer with id:%s", id),
		})
		return
	}

	var input models.EditCustomerInput
	if err := context.BindJSON(&input); err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config.DB.Model(&customer).Updates(input)

	context.JSON(http.StatusOK, gin.H{"data": customer})
}

// GetAllCustomers  Retrieve all customers
func GetAllCustomers(context *gin.Context) {
	var customers []models.Customer
	var err = config.DB.
		Preload("Addresses").
		Preload("Addresses.Coordinate").
		Preload("Contacts").
		Find(&customers).Error

	if err != nil {
		context.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": customers})
}

// GetACustomer find customer by id param
func GetACustomer(context *gin.Context) {
	var customer models.Customer
	var id = context.Param("id")
	var err = config.DB.
		Preload("Addresses").
		Preload("Addresses.Coordinate").
		Preload("Contacts").
		Where("id = ?", id).
		First(&customer).Error

	if err != nil {
		context.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"message": fmt.Sprintf("Cannot find customer with id:%s", id),
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": customer})
}

// DeleteCustomer delete customer by id param
func DeleteCustomer(context *gin.Context) {
	var customer models.Customer
	var id = context.Param("id")

	if err := config.DB.Where("id = ?", id).First(&customer).Error; err != nil {
		context.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"message": fmt.Sprintf("Cannot find customer with id:%s", id),
		})
		return
	}

	config.DB.Delete(&customer)

	context.JSON(http.StatusOK, gin.H{"data": true})
}
