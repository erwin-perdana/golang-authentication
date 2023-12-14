package controller

import (
	"time"
	"net/http"
	"github.com/gin-gonic/gin"
	"golang-authentication/initializers"
	"golang-authentication/models"
)

func Pay(c *gin.Context) {
	var body struct {
		TransactionID uint
	}

	var transaction models.Transaction

	if err := c.Bind(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to process request data",
		})
		return
	}

	if err := initializers.DB.First(&transaction, body.TransactionID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Transaction not found",
		})
		return
	}

	payment := models.Payment{
		TransactionID: body.TransactionID,
		Date: time.Now(),
		Amount: transaction.Nominal,
		Method: "Transfer",
		Status: "Selesai",
	}

	result := initializers.DB.Create(&payment)

	transaction.Status = "Selesai"

	if err := initializers.DB.Save(&transaction).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed update transaction status",
		})
		
		return
	}

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create payment",
		})

		return
	}

	CreateHistory(c, "Pay", 0)

	c.JSON(http.StatusOK, gin.H{
		"message": "Success create payment",
	})
}
