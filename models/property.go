package models

import "github.com/google/uuid"

type PropertyType string

const (
	Apartment  PropertyType = "apartment"
	House      PropertyType = "house"
	Commercial PropertyType = "commercial"
	Industrial PropertyType = "industrial"
)

type Property struct {
	ID uuid.UUID `json:"id" gorm:"primaryKey"`
}
