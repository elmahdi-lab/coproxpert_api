package models

import (
	"github.com/google/uuid"
)

type Role string
type EntityType string

const (
	PropertyEntity     EntityType = "property"
	BuildingEntity     EntityType = "building"
	OrganizationEntity EntityType = "organization"
)

const (
	AdminRole   Role = "admin_role"
	ManagerRole Role = "manager_role"
	UserRole    Role = "user_role"
)

type Permission struct {
	ID         uuid.UUID `json:"id" gorm:"type:uuid;primaryKey"`
	UserID     uuid.UUID
	User       *User       `json:"user" gorm:"foreignKey:UserID;references:ID;constraint:OnDelete:CASCADE"`
	Role       *Role       `json:"role" gorm:"not null;check:role IN ('admin_role', 'manager_role', 'user_role')"`
	EntityType *EntityType `json:"entity_type" gorm:"not null;check:entity_type IN ('property', 'building', 'organization')"`
	EntityID   *uuid.UUID  `json:"entity_id" gorm:"type:uuid"`
	BaseModel
}

/*
id, user_id, access_level, entity, entity_id, created_at, updated_at
admin_role, nil, nil
manager_role, organization, id
manager_role, building, id
*/
