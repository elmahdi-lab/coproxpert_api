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
	ID   uuid.UUID    `json:"id" gorm:"type:uuid;primaryKey"`
	Name string       `json:"name" gorm:"unique"`
	Type PropertyType `json:"type"`

	BuildingID uuid.UUID  `json:"building_id" gorm:"type:uuid"`
	Building   *UnitGroup `json:"building" gorm:"foreignKey:BuildingID;references:ID;constraint:OnDelete:CASCADE"`

	UserID uuid.UUID `json:"user_id" gorm:"type:uuid"`
	User   *User     `json:"user" gorm:"foreignKey:UserID;references:ID;constraint:OnDelete:CASCADE"`

	IsEnabled *bool `json:"is_enabled"`

	BaseModel
}
