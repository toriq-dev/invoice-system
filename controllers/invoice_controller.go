package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"invoice-system/database"
	"invoice-system/models"
)

func GetInvoices(c *gin.Context) {
	var invoices []models.Invoice
	database.DB.Preload("InvoiceItems").Find(&invoices)
	c.JSON(http.StatusOK, invoices)
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
