package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	TenantID     *uuid.UUID `gorm:"type:uuid"`
	ID           *uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	FirstName    string     `gorm:"size:50;not null"`
	MiddleName   string     `gorm:"size:50"`
	LastName     string     `gorm:"size:50;not null"`
	Email        string     `gorm:"size:100;uniqueIndex;not null"`
	PasswordHash string     `gorm:"size:100;not null"`
	Roles        []Role     `gorm:"many2many:user_roles"`
	CreatedAt    time.Time  `gorm:"not null;default:now()"`
	UpdatedAt    time.Time  `gorm:"not null;default:now()"`
}
