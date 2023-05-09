package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	TenantID     uuid.UUID
	ID           uuid.UUID
	FirstName    string
	MiddleName   string
	LastName     string
	Email        string
	PasswordHash string
	Roles        []Role
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
