package initializers

import "github.com/vaibhavsingh9/go-jwt/models"

func SyncDatabase() {
	DB.AutoMigrate(&models.User{})
}
