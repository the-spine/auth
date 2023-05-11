package repository

import (
	"auth/internal/db"
	"auth/internal/models"
)

func CreateUser(user *models.User) error {
	result := db.DB.Create(user)
	return result.Error
}
