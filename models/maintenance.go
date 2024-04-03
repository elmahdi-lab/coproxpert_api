package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type MaintenanceType string

const (
	PlumbingMaintenance   MaintenanceType = "Plumbing"
	ElectricalMaintenance MaintenanceType = "Electrical"
	GeneralMaintenance    MaintenanceType = "General"
	GardeningMaintenance  MaintenanceType = "Gardening"
	PoolMaintenance       MaintenanceType = "Pool"
	PaintingMaintenance   MaintenanceType = "Painting"
	Exterminator          MaintenanceType = "Exterminator"
)

type Maintenance struct {
	ID          uuid.UUID       `json:"id" gorm:"type:uuid; primaryKey"`
	UnitGroupId uuid.UUID       `json:"unit_group_id" gorm:"type:uuid; not null"`
	UnitId      uuid.UUID       `json:"unit_id" gorm:"type:uuid; not null"`
	Type        MaintenanceType `json:"type" gorm:"not null;"`
	Comment     string          `json:"comment" gorm:"type:varchar(255)"`
	StartDate   time.Time       `json:"start_date" gorm:"not null"`
	EndDate     time.Time       `json:"end_date" gorm:"not null; check:end_date > start_date"`
	IsDone      bool            `json:"is_done" gorm:"default:false; not null"`
	BaseModel
}

// BeforeCreate hook to set default values and perform validation
func (m *Maintenance) BeforeCreate(tx *gorm.DB) (err error) {
	// Perform custom validation
	if m.StartDate.After(m.EndDate) {
		return gorm.ErrInvalidData
	}
	return nil
}
