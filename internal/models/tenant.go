package models

import (
	"time"

	"github.com/google/uuid"
)

type Tenant struct {
	ID          uuid.UUID
	Name        string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
