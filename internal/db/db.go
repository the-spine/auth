package db

import (
	"auth/internal/config"
	"auth/internal/models"
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func ConnectDB(config *config.Config) error {

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  logger.Info,
			IgnoreRecordNotFoundError: true,
			ParameterizedQueries:      true,
			Colorful:                  true,
		},
	)

	var err error

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable", config.Postgres.Host, config.Postgres.User, config.Postgres.Password, config.Postgres.DBName, config.Postgres.Port)

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{Logger: newLogger})

	if err != nil {
		return err
	}

	if err := DB.Exec(`CREATE EXTENSION IF NOT EXISTS "uuid-ossp";`).Error; err != nil {
		return err
	}

	log.Println("Running Migrations")

	err = DB.AutoMigrate(&models.User{}, &models.Role{}, &models.Tenant{})

	return err
}
