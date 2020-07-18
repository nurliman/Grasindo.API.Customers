package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/nurliman/Grasindo.API.go/config"
	"github.com/nurliman/Grasindo.API.go/models"

	"github.com/gin-gonic/gin"
)

// AddContactToCustomer adding a contact to a customer
func AddContactToCustomer(context *gin.Context) {
	var customer models.Customer
	var input models.Contact
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
		Association("Contacts").
		Append(&input).
		Error; err != nil {
		context.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": input})
}

// EditCustomerContact edit a contact of a customer
func EditCustomerContact(context *gin.Context) {
	var customer models.Customer
	var contact models.Contact
	var customerID = context.Param("customerId")
	var contactID = context.Param("contactId")

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
		Where("id = ? AND customer_id = ?", contactID, customerID).
		First(&contact).
		Error; err != nil {
		context.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"message": fmt.Sprintf(
				"Cannot find contact with id:%s in custumer with id:%s",
				contactID,
				customerID,
			),
		})
		return
	}

	var input models.ContactInput
	if err := context.BindJSON(&input); err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config.DB.Model(&contact).Update(&input)

	context.JSON(http.StatusOK, gin.H{"data": contact})
}

// GetCustomerContacts get contacts of a customer
func GetCustomerContacts(context *gin.Context) {
	var customer models.Customer
	var customerID = context.Param("customerId")

	if err := config.DB.
		Preload("Contacts").
		Select("id").
		Where("id = ?", customerID).
		First(&customer).Error; err != nil {
		context.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"message": fmt.Sprintf("Cannot find customer with id:%s", customerID),
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": customer.Contacts})
}

// GetCustomerContact find customer's contact by id
func GetCustomerContact(context *gin.Context) {

	var customer models.Customer
	var customerID = context.Param("customerId")
	var contactID, contactIDError = strconv.Atoi(context.Param("contactId"))

	if contactIDError != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "contact id must number",
		})
		return
	}

	if err := config.DB.
		Preload("Contacts").
		Select("id").
		Where("id = ?", customerID).
		First(&customer).Error; err != nil {
		context.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"message": fmt.Sprintf("Cannot find customer with id:%s", customerID),
		})
		return
	}

	for index, element := range customer.Contacts {
		if element.ID == uint(contactID) {
			context.JSON(http.StatusOK, gin.H{"data": customer.Contacts[index]})
			return
		}
	}

	context.AbortWithStatusJSON(http.StatusNotFound, gin.H{
		"message": fmt.Sprintf(
			"Cannot find contact with id:%d in customer with id:%s",
			contactID,
			customerID,
		),
	})
}

// DeleteCustomerContact delete a contact from customer
func DeleteCustomerContact(context *gin.Context) {
	var customer models.Customer
	var contact models.Contact
	var customerID = context.Param("customerId")
	var contactID = context.Param("contactId")

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
		Where("id = ? AND customer_id = ?", contactID, customerID).
		First(&contact).
		Error; err != nil {
		context.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"message": fmt.Sprintf(
				"Cannot find address with id:%s in custumer with id:%s",
				contactID,
				customerID,
			),
		})
		return
	}

	config.DB.Delete(&contact)

	context.JSON(http.StatusOK, gin.H{"data": true})
}
