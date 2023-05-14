package models

import (
	"time"

	"github.com/google/uuid"
)

type Tenant struct {
	ID          uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Name        string    `gorm:"size:50;uniqueIndex;not null"`
	Description string    `gorm:"size:255"`
	CreatedAt   time.Time `gorm:"not null"`
	UpdatedAt   time.Time `gorm:"not null"`
}
