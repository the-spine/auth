package models

import "github.com/google/uuid"

type Role struct {
	TenantID    uuid.UUID
	ID          uuid.UUID
	Role        string
	Permissions []string
}
