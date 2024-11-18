package routes

import (
	"invoice-system/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	api := router.Group("/api")
	{
		api.GET("/invoices", controllers.GetInvoices)
		api.GET("/invoices/:id", controllers.GetInvoiceByID)
		api.POST("/invoices", controllers.CreateInvoice)
		api.PUT("/invoices/:id", controllers.UpdateInvoice)
		api.DELETE("/invoices/:id", controllers.DeleteInvoice)
	}
}
