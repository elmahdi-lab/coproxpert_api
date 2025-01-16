package models

import (
	"github.com/google/uuid"
)

type PropertyType string

const (
	Apartment PropertyType = "a"
	House     PropertyType = "h"
	Villa     PropertyType = "v"
)

type Unit struct {
	ID uuid.UUID `json:"id" gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`

	UserID *uuid.UUID `json:"user_id" gorm:"type:uuid;constraint:OnDelete:CASCADE;"`
	User   *User      `json:"user" gorm:"foreignKey:UserID;references:ID;constraint:OnDelete:CASCADE"`

	Name string       `json:"name" gorm:"not null"`
	Type PropertyType `json:"type" gorm:"default:'a'"`

	UnitGroupID uuid.UUID  `json:"unit_group_id" gorm:"type:uuid"`
	UnitGroup   *UnitGroup `json:"unit_group" gorm:"foreignKey:UnitGroupID;references:ID;constraint:OnDelete:CASCADE"`
	IsEnabled   *bool      `json:"is_enabled" gorm:"default:true"`

	BaseModel
}
