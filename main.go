package main

import (
	"github.com/gin-gonic/gin"
	"invoice-system/database"
	"invoice-system/models"
	"invoice-system/routes"
)

func main() {
	// Connect to database
	database.Connect()

	// Migrate models
	database.DB.AutoMigrate(&models.Invoice{}, &models.InvoiceItem{})

	// Setup Gin router
	router := gin.Default()
	routes.SetupRoutes(router)

	// Start server
	router.Run(":8080")
}
