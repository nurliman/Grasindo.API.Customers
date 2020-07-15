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

		// endpoints "/v1/customers"
		customers := v1.Group("/customers")
		{
			customers.GET("/", controllers.GetAllCustomers)
			customers.GET("/:id", controllers.GetACustomer)
			customers.POST("/", controllers.AddCustomer)
			customers.PUT("/:id", controllers.AddCustomer)
			customers.DELETE("/:id", controllers.DeleteACustomer)
		}
	}

	return router
}
