package models

import (
	"github.com/google/uuid"
)

type UnitGroup struct {
	ID             uuid.UUID      `json:"id" gorm:"type:uuid;primaryKey"`
	Name           string         `json:"name" gorm:"not null"`
	OrganizationID uuid.UUID      `json:"organization_id" gorm:"type:uuid"`
	UserID         uuid.UUID      `json:"user_id" gorm:"type:uuid"`
	Maintenances   *[]Maintenance `json:"maintenances" gorm:"foreignKey:UnitGroupId;references:ID;constraint:OnDelete:CASCADE"`
	BaseModel
}
