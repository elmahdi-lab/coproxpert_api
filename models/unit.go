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
	ID     uuid.UUID    `json:"id" gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	UserID *uuid.UUID   `json:"userID" gorm:"type:uuid;constraint:OnDelete:CASCADE;"`
	User   *User        `json:"user" gorm:"foreignKey:UserID;references:ID;constraint:OnDelete:CASCADE"`
	Name   string       `json:"name" gorm:"unique"`
	Type   PropertyType `json:"type"`

	UnitGroupID uuid.UUID `json:"unit_group_id" gorm:"type:uuid"`
	IsEnabled   *bool     `json:"is_enabled"`

	BaseModel
}
