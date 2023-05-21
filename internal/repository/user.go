package repository

import (
	"auth/internal/db"
	"auth/internal/models"
)

func CreateUser(user *models.User) error {
	result := db.DB.Create(user)
	return result.Error
}

func GetUserByEmail(email string, user *models.User) error {
	result := db.DB.Where("email = ?").First(user)
	return result.Error
}

func GetUserById(id string, user *models.User) error {
	result := db.DB.Where("id = ?").First(id)
	return result.Error
}
