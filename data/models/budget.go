package models

import "github.com/google/uuid"

// Budget is a model for budgets, they may be provisional or definitive
type Budget struct {
	ID            uuid.UUID `json:"id" gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	OwnerID       uuid.UUID `json:"owner_id" gorm:"type:uuid;not null"`
	UnitGroupID   uuid.UUID `json:"unit_group_id" gorm:"type:uuid;not null"`
	Amount        float64   `json:"amount" gorm:"not null"`
	Year          int       `json:"year" gorm:"not null"`
	IsProvisional bool      `json:"is_provisional" gorm:"default:true"`
	BaseModel
}
