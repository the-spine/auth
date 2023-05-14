package db

import (
	"auth/internal/config"
	"auth/internal/models"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB(config *config.Config) error {
	var err error
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable", config.Postgres.Host, config.Postgres.User, config.Postgres.Password, config.Postgres.DBName, config.Postgres.Port)
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return err
	}

	DB.Raw("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\";")

	DB.AutoMigrate(&models.User{}, &models.Role{}, &models.Tenant{})

	return err
}
