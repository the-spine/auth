package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	TenantID     uuid.UUID `gorm:"type:uuid;not null"`
	ID           uuid.UUID `gorm:"type:uuid;primaryKey"`
	FirstName    string    `gorm:"size:50;not null"`
	MiddleName   string    `gorm:"size:50"`
	LastName     string    `gorm:"size:50;not null"`
	Email        string    `gorm:"size:100;uniqueIndex;not null"`
	PasswordHash string    `gorm:"size:100;not null"`
	Roles        []Role    `gorm:"many2many:user_roles"`
	CreatedAt    time.Time `gorm:"not null"`
	UpdatedAt    time.Time `gorm:"not null"`
}
