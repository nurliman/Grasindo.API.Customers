package controllers

import (
	"fmt"
	"net/http"

	"github.com/nurliman/Grasindo.API.Customers/config"
	"github.com/nurliman/Grasindo.API.Customers/models"

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

// EditCustomer update edit customer information, but not updating sub models like Address, Customer
func EditCustomer(context *gin.Context) {
	var customer models.Customer
	var customerID = context.Param("customerId")

	if err := config.DB.
		Preload("Addresses").
		Preload("Addresses.Coordinate").
		Preload("Contacts").
		Where("id = ?", customerID).
		First(&customer).
		Error; err != nil {
		context.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"message": fmt.Sprintf("Cannot find customer with id:%s", customerID),
		})
		return
	}

	var input models.EditCustomerInput
	if err := context.BindJSON(&input); err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config.DB.Model(&customer).Set("gorm:association_save_reference", false).Update(&input)

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

// GetCustomer find customer by id param
func GetCustomer(context *gin.Context) {
	var customer models.Customer
	var customerID = context.Param("customerId")
	var err = config.DB.
		Preload("Addresses").
		Preload("Addresses.Coordinate").
		Preload("Contacts").
		Where("id = ?", customerID).
		First(&customer).Error

	if err != nil {
		context.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"message": fmt.Sprintf("Cannot find customer with id:%s", customerID),
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": customer})
}

// DeleteCustomer delete customer by id param
func DeleteCustomer(context *gin.Context) {
	var customer models.Customer
	var customerID = context.Param("customerId")

	if err := config.DB.Where("id = ?", customerID).First(&customer).Error; err != nil {
		context.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"message": fmt.Sprintf("Cannot find customer with id:%s", customerID),
		})
		return
	}

	config.DB.Delete(&customer)

	context.JSON(http.StatusOK, gin.H{"data": true})
}
