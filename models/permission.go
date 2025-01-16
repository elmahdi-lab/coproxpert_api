package models

import (
	"github.com/google/uuid"
)

type Role string
type EntityType string

const (
	SuperAdminRole Role = "sa"
	AdminRole      Role = "a"
	ManagerRole    Role = "m"
	UserRole       Role = "u" // TODO: this may not be necessary
)

const (
	UnitGroupEntity EntityType = "unit_group"
	UnitEntity      EntityType = "unit"
)

type Permission struct {
	ID         string     `json:"id" gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	UserID     uuid.UUID  `json:"user_id" gorm:"type:uuid"`
	EntityID   uuid.UUID  `json:"entity_id" gorm:"type:uuid"`
	EntityType EntityType `json:"entity_type" gorm:"type:uuid"`
	Role       Role       `json:"role" gorm:"not null; default:'u'"`
	BaseModel
}

func (p *Permission) IsAdmin() bool {
	return p.Role == AdminRole
}
