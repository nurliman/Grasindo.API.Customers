package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/nurliman/Grasindo.API.Customer/config"
	"github.com/nurliman/Grasindo.API.Customer/models"

	"github.com/gin-gonic/gin"
)

// AddAddressToCustomer adding an address to a customer
func AddAddressToCustomer(context *gin.Context) {
	var customer models.Customer
	var input models.Address
	var customerID = context.Param("customerId")

	if err := config.DB.
		Select("id").
		Where("id = ?", customerID).
		First(&customer).
		Error; err != nil {
		context.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"message": fmt.Sprintf("Cannot find customer with id:%s", customerID),
		})
		return
	}

	if err := context.BindJSON(&input); err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := config.DB.
		Model(&customer).
		Association("Addresses").
		Append(&input).
		Error; err != nil {
		context.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": input})
}

// EditCustomerAddress edit an address of a customer
func EditCustomerAddress(context *gin.Context) {
	var customer models.Customer
	var address models.Address
	var customerID = context.Param("customerId")
	var addressID = context.Param("addressId")

	if err := config.DB.
		Select("id").
		Where("id = ?", customerID).
		First(&customer).
		Error; err != nil {
		context.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"message": fmt.Sprintf("Cannot find customer with id:%s", customerID),
		})
		return
	}

	if err := config.DB.
		Preload("Coordinate").
		Where("id = ? AND customer_id = ?", addressID, customerID).
		First(&address).
		Error; err != nil {
		context.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"message": fmt.Sprintf(
				"Cannot find address with id:%s in custumer with id:%s",
				addressID,
				customerID,
			),
		})
		return
	}

	var input models.AddressInput
	input.Coordinate = address.Coordinate
	if err := context.BindJSON(&input); err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config.DB.Model(&address).Update(&input)

	context.JSON(http.StatusOK, gin.H{"data": address})
}

// GetCustomerAddresses retrieve all addresses in a customer data
func GetCustomerAddresses(context *gin.Context) {
	var customer models.Customer
	var customerID = context.Param("customerId")

	if err := config.DB.
		Preload("Addresses").
		Preload("Addresses.Coordinate").
		Select("id").
		Where("id = ?", customerID).
		First(&customer).Error; err != nil {
		context.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"message": fmt.Sprintf("Cannot find customer with id:%s", customerID),
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": customer.Addresses})
}

// GetCustomerAddress find customer's address by id
func GetCustomerAddress(context *gin.Context) {

	var customer models.Customer
	var customerID = context.Param("customerId")
	var addressID, addressIDError = strconv.Atoi(context.Param("addressId"))

	if addressIDError != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "address id must number",
		})
		return
	}

	if err := config.DB.
		Preload("Addresses").
		Preload("Addresses.Coordinate").
		Select("id").
		Where("id = ?", customerID).
		First(&customer).Error; err != nil {
		context.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"message": fmt.Sprintf("Cannot find customer with id:%s", customerID),
		})
		return
	}

	for index, element := range customer.Addresses {
		if element.ID == uint(addressID) {
			context.JSON(http.StatusOK, gin.H{"data": customer.Addresses[index]})
			return
		}
	}

	context.AbortWithStatusJSON(http.StatusNotFound, gin.H{
		"message": fmt.Sprintf(
			"Cannot find address with id:%d in customer with id:%s",
			addressID,
			customerID,
		),
	})
}

// DeleteCustomerAddress delete an address from customer
func DeleteCustomerAddress(context *gin.Context) {
	var customer models.Customer
	var address models.Address
	var customerID = context.Param("customerId")
	var addressID = context.Param("addressId")

	if err := config.DB.
		Select("id").
		Where("id = ?", customerID).
		First(&customer).
		Error; err != nil {
		context.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"message": fmt.Sprintf("Cannot find customer with id:%s", customerID),
		})
		return
	}

	if err := config.DB.
		Preload("Coordinate").
		Where("id = ? AND customer_id = ?", addressID, customerID).
		First(&address).
		Error; err != nil {
		context.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"message": fmt.Sprintf(
				"Cannot find address with id:%s in custumer with id:%s",
				addressID,
				customerID,
			),
		})
		return
	}

	config.DB.Delete(&address.Coordinate)
	config.DB.Delete(&address)

	context.JSON(http.StatusOK, gin.H{"data": true})
}
