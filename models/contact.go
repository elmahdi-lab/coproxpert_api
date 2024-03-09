package models

import (
	"github.com/google/uuid"
)

type Contact struct {
	ID          uuid.UUID `json:"id" gorm:"primaryKey"`
	UserID      uuid.UUID
	User        *User   `json:"user" gorm:"foreignKey:UserID;references:ID;constraint:OnDelete:CASCADE"`
	IsDefault   *bool   `json:"is_default"`
	PhoneNumber *string `json:"phone_number"`
	Address     *string `json:"address"`
	City        *string `json:"city"`
	Province    *string `json:"province"`
	ZipCode     *string `json:"zip_code"`
	Country     *string `json:"country"`
}
