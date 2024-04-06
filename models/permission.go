package models

import (
	"github.com/google/uuid"
)

type Role string

const (
	AdminRole   Role = "admin_role"
	ManagerRole Role = "manager_role"
)

type Permission struct {
	ID     string    `json:"id" gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	UserID uuid.UUID `json:"user_id" gorm:"type:uuid"`
	Role   Role      `json:"role" gorm:"not null"`

	BaseModel
}

// (userID, 'manager_role', 'organization', 'organization_id')
// (yassine, manager, via_syndic, via_syndic_id)
// (tajeddine, manager, AUTO, AUTO_ID)
