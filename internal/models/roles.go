package models

import (
	"github.com/google/uuid"
	authpb "github.com/the-spine/spine-protos-go/auth"
)

type Role struct {
	TenantID    uuid.UUID `gorm:"type:uuid;not null"`
	ID          uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Role        string    `gorm:"size:50;uniqueIndex;not null"`
	Permissions []string  `gorm:"type:text[]"`
}

func (r *Role) ToRespone() authpb.Role {

	return authpb.Role{
		Role:        r.Role,
		Permissions: r.Permissions,
	}
}
