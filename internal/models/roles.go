package models

import "github.com/google/uuid"

type Role struct {
	TenantID    uuid.UUID `gorm:"type:uuid;not null"`
	ID          uuid.UUID `gorm:"type:uuid;primaryKey"`
	Role        string    `gorm:"size:50;uniqueIndex;not null"`
	Permissions []string  `gorm:"type:text[]"`
}
