package models

import "github.com/google/uuid"

type Credential struct {
	ID     int `json:"id" gorm:"primaryKey"`
	UserID uuid.UUID
	User   *User `json:"user" gorm:"foreignKey:UserID;references:ID;constraint:OnDelete:CASCADE"`
}
