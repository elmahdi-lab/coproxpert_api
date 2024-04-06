package models

import (
	"github.com/google/uuid"
)

type PropertyType string

const (
	Apartment PropertyType = "apartment"
	House     PropertyType = "house"
	Villa     PropertyType = "villa"
)

type Unit struct {
	ID uuid.UUID `json:"id" gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`

	Name string       `json:"name" gorm:"unique"`
	Type PropertyType `json:"type"`

	BuildingID uuid.UUID `json:"building_id" gorm:"type:uuid"`
	UserID     uuid.UUID `json:"user_id" gorm:"type:uuid"`

	IsEnabled *bool `json:"is_enabled"`

	BaseModel
}
