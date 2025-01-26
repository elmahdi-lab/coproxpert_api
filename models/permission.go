package models

import (
	"github.com/google/uuid"
)

type Role int8
type EntityName string

const (
	UnitGroupEntity EntityName = "unit_group"
	UnitEntity      EntityName = "unit"
)

const (
	SuperAdminRole Role = 127
	AdminRole      Role = 100
	ManagerRole    Role = 50
	UserRole       Role = 10
)

type Permission struct {
	ID         string     `json:"id" gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	UserID     uuid.UUID  `json:"user_id" gorm:"type:uuid"`
	EntityID   uuid.UUID  `json:"entity_id" gorm:"type:uuid"`
	EntityName EntityName `json:"entity_name" gorm:"not null"`
	Role       Role       `json:"role" gorm:"not null; default:10"`
	BaseModel
}

func (p *Permission) IsAdmin() bool {
	return p.Role == AdminRole
}
