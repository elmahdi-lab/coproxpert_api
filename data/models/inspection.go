package models

import "time"

type Inspection struct {
	ID           string    `json:"id" gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	UnitGroupID  string    `json:"unit_group_id" gorm:"type:uuid;not null"`
	OwnerID      string    `json:"owner_id" gorm:"type:uuid;not null"`
	AssignedTo   string    `json:"assigned_to" gorm:"type:uuid;not null"`
	StartDate    string    `json:"start_date" gorm:"type:date;not null"`
	EndDate      string    `json:"end_date" gorm:"type:date;not null"`
	IsComplete   bool      `json:"active" gorm:"default:false"`
	IsCompleteAt time.Time `json:"is_complete_at" gorm:"type:timestamp"`
	Details      string    `json:"details" gorm:"type:text"`
	BaseModel
}
