package models

import (
	"github.com/google/uuid"
)

type Contact struct {
	ID          uuid.UUID `json:"id" gorm:"type:uuid;primaryKey"`
	UserID      uuid.UUID `json:"userID" gorm:"type:uuid"`
	User        *User     `json:"user" gorm:"foreignKey:UserID;references:ID;constraint:OnDelete:CASCADE"`
	PhoneNumber *string   `json:"phone_number"`
	Address     *string   `json:"address"`
	City        *string   `json:"city"`
	Province    *string   `json:"province"`
	ZipCode     *string   `json:"zip_code"`
	Country     *string   `json:"country"`
	BaseModel
}
