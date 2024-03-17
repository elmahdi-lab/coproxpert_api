package models

import (
	"github.com/google/uuid"
	"time"
)

type Maintenance struct {
	ID         uuid.UUID `json:"id" gorm:"type:uuid; primaryKey"`
	BuildingID uuid.UUID `json:"building_id" gorm:"type:uuid"`
	Building   *Building `json:"building" gorm:"foreignKey:BuildingID;references:ID;constraint:OnDelete:CASCADE"`
	PropertyID uuid.UUID `json:"property_id" gorm:"type:uuid"`
	Property   *Property `json:"property" gorm:"foreignKey:PropertyID;references:ID;constraint:OnDelete:CASCADE"`
	StartDate  time.Time `json:"start_date"`
	EndDate    time.Time `json:"end_date"`
	IsDone     bool      `json:"is_done" gorm:"default:false"`
	BaseModel
}
