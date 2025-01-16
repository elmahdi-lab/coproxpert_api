package models

import (
	"github.com/google/uuid"
)

type UnitGroup struct {
	ID uuid.UUID `json:"id" gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`

	Name string `json:"name" gorm:"not null"`

	UserID *uuid.UUID `json:"userID" gorm:"type:uuid;constraint:OnDelete:CASCADE;"`
	User   *User      `json:"user" gorm:"foreignKey:UserID;references:ID;constraint:OnDelete:CASCADE"`

	BaseModel
}
