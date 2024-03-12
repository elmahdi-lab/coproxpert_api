package models

import (
	"github.com/google/uuid"
)

type PropertyType string

const (
	Apartment  PropertyType = "apartment"
	House      PropertyType = "house"
	Commercial PropertyType = "commercial"
	Industrial PropertyType = "industrial"
)

type Property struct {
	ID   uuid.UUID    `json:"id" gorm:"type:uuid;primaryKey"`
	Name string       `json:"name" gorm:"unique"`
	Type PropertyType `json:"type"`

	BuildingID uuid.UUID `json:"building_id" gorm:"type:uuid"`
	Building   *Building `json:"building" gorm:"foreignKey:BuildingID;references:ID;constraint:OnDelete:CASCADE"`

	IsEnabled *bool `json:"is_enabled"`

	BaseModel
}
