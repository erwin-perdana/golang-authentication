package controller

import (
	"time"
	"net/http"
	"github.com/gin-gonic/gin"
	"golang-authentication/initializers"
	"golang-authentication/models"
)

func CreateHistory(c *gin.Context, action string, userID uint) {
	if userID == 0 {
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
	
		userID = user.ID
	}

	history := models.History{
		UserID:  userID,
		Date:    time.Now(),
		Action: action,
	}

	result := initializers.DB.Create(&history)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create history",
		})

		return
	}
}
