package models

import (
	"github.com/google/uuid"
)

type AccessLevel string
type EntityType string

const (
	PropertyEntity     EntityType = "property"
	BuildingEntity     EntityType = "building"
	OrganizationEntity EntityType = "organization"
)

const (
	AdminRole AccessLevel = "admin_role"
	UserRole  AccessLevel = "user_role"
)

type Permission struct {
	ID          uuid.UUID `json:"id" gorm:"type:uuid;primaryKey"`
	UserID      uuid.UUID
	User        *User        `json:"user" gorm:"foreignKey:UserID;references:ID;constraint:OnDelete:CASCADE"`
	AccessLevel *AccessLevel `json:"access_level" gorm:"not null;check:access_level IN ('admin_role', 'user_role')"`
	Entity      *EntityType  `json:"entity" gorm:"not null;check:entity IN ('property', 'building', 'organization')"`
	EntityID    *uuid.UUID   `json:"entity_id"`
	BaseModel
}
