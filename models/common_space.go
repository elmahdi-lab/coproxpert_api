package models

import (
	"github.com/google/uuid"
)

type CommonSpace struct {
	ID         uuid.UUID `json:"id" gorm:"type:uuid;primaryKey"`
	BuildingID uuid.UUID
	Building   *Building `json:"building" gorm:"foreignKey:BuildingID;references:ID;constraint:OnDelete:CASCADE"`
	BaseModel
}
