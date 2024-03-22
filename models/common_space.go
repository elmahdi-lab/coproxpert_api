package models

import (
	"github.com/google/uuid"
)

type CommonSpace struct {
	ID         uuid.UUID  `json:"id" gorm:"type:uuid;primaryKey"`
	BuildingID uuid.UUID  `json:"building_id" gorm:"type:uuid"`
	Building   *UnitGroup `json:"building" gorm:"foreignKey:BuildingID;references:ID;constraint:OnDelete:CASCADE"`
	BaseModel
}
