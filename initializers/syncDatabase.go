package initializers

import "golang-authentication/models"

func SyncDatabase() {
	DB.AutoMigrate(&models.User{}, &models.Transaction{}, &models.Payment{})
}