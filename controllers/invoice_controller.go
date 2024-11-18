package controllers

import (
	"net/http"
	"strconv"

	"invoice-system/database"
	"invoice-system/models"

	"github.com/gin-gonic/gin"
)

func GetInvoices(c *gin.Context) {
	var invoices []models.Invoice

	// Default values for pagination
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 10
	}

	// Calculate offset
	offset := (page - 1) * limit

	// Query with pagination
	result := database.DB.Preload("InvoiceItems").
		Limit(limit).
		Offset(offset).
		Find(&invoices)

	// Check for database errors
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to retrieve invoices",
			"details": result.Error.Error(),
		})
		return
	}

	// Get total count for pagination metadata
	var total int64
	database.DB.Model(&models.Invoice{}).Count(&total)

	// Build response with pagination metadata
	c.JSON(http.StatusOK, gin.H{
		"data": invoices,
		"pagination": gin.H{
			"current_page": page,
			"page_size":    limit,
			"total_pages":  (total + int64(limit) - 1) / int64(limit), // Calculate total pages
			"total_items":  total,
		},
	})
}

func GetInvoiceByID(c *gin.Context) {
	var invoice models.Invoice
	id := c.Param("id")

	if err := database.DB.Preload("InvoiceItems").First(&invoice, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Invoice not found"})
		return
	}

	c.JSON(http.StatusOK, invoice)
}

func CreateInvoice(c *gin.Context) {
	var invoice models.Invoice

	if err := c.ShouldBindJSON(&invoice); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.DB.Create(&invoice)
	c.JSON(http.StatusCreated, invoice)
}

func UpdateInvoice(c *gin.Context) {
	var invoice models.Invoice
	id := c.Param("id")

	if err := database.DB.First(&invoice, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Invoice not found"})
		return
	}

	if err := c.ShouldBindJSON(&invoice); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.DB.Save(&invoice)
	c.JSON(http.StatusOK, invoice)
}

func DeleteInvoice(c *gin.Context) {
	var invoice models.Invoice
	id := c.Param("id")

	if err := database.DB.First(&invoice, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Invoice not found"})
		return
	}

	database.DB.Delete(&invoice)
	c.JSON(http.StatusOK, gin.H{"message": "Invoice deleted"})
}
