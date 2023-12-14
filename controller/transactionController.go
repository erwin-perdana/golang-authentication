package controller

import (
	"time"
	"net/http"
	"github.com/gin-gonic/gin"
	"golang-authentication/initializers"
	"golang-authentication/models"
)

func Buy(c *gin.Context) {
	var body struct {
		Nominal int
	}

	if err := c.Bind(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to process request data",
		})
		return
	}

	userInterface, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "User not found",
		})
	}

	user, ok := userInterface.(models.User)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Unexpected user type",
		})
		return
	}

	userID := user.ID

	transaction := models.Transaction{
		UserID:  userID,
		Date:    time.Now(),
		Status:  "Pending",
		Nominal: body.Nominal,
	}

	result := initializers.DB.Create(&transaction)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create transaction",
		})

		return
	}

	CreateHistory(c, "Buy", 0)

	c.JSON(http.StatusOK, gin.H{
		"message": "Success create transaction",
	})
}

func GetTransactions(c *gin.Context) {
	transactionID := c.Param("id")

	var transaction models.Transaction

	if err := initializers.DB.First(&transaction, transactionID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Transaction not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"transaction": transaction,
	})
}
