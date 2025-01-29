package models

import (
	"github.com/google/uuid"
)

type UnitGroup struct {
	ID      uuid.UUID `json:"id" gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	Name    string    `json:"name" gorm:"not null"`
	OwnerID uuid.UUID `json:"owner_id" gorm:"type:uuid"`
	Owner   *User     `json:"owner" gorm:"foreignKey:OwnerID;references:ID;constraint:OnDelete:CASCADE"`
	BaseModel
}
