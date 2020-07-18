package routes

import (
	"net/http"

	"github.com/nurliman/Grasindo.API.go/controllers"

	"github.com/gin-gonic/gin"
)

// SetupRouter function define routes endpoints
func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{"api": "grasindo"})
	})

	// endpoints "/v1"
	v1 := router.Group("/v1")
	{
		v1.GET("/", func(context *gin.Context) {
			context.JSON(http.StatusOK, gin.H{"api": "grasindo", "version": "1"})
		})

		// customers endpoints "/v1/customers"
		customers := v1.Group("/customers")
		{
			customers.GET("/", controllers.GetAllCustomers)
			customers.GET("/:customerId", controllers.GetCustomer)
			customers.POST("/", controllers.AddCustomer)
			customers.PUT("/:customerId", controllers.EditCustomer)
			customers.DELETE("/:customerId", controllers.DeleteCustomer)

			addresses := customers.Group("/:customerId/addresses")
			{
				addresses.GET("/", controllers.GetCustomerAddresses)
				addresses.GET("/:addressId", controllers.GetCustomerAddress)
				addresses.POST("/", controllers.AddAddressToCustomer)
				addresses.PUT("/:addressId", controllers.EditCustomerAddress)
				addresses.DELETE("/:addressId", controllers.DeleteCustomerAddress)
			}

			contacts := customers.Group("/:customerId/contacts")
			{
				contacts.GET("/", controllers.GetCustomerContacts)
				contacts.GET("/:contactId", controllers.GetCustomerContact)
				contacts.POST("/", controllers.AddContactToCustomer)
				contacts.DELETE("/:contactId", controllers.DeleteCustomerContact)
			}
		}
	}

	return router
}
