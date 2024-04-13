package models

import "github.com/google/uuid"

// Budget is a model for budgets, they may be provisional or definitive
type Budget struct {
	ID             uuid.UUID `json:"id" gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	OrganizationID uuid.UUID `json:"organization_id" gorm:"type:uuid"`
	Year           int       `json:"year" gorm:"not null"`
	IsProvisional  bool      `json:"is_provisional" gorm:"default:true"`
	BaseModel
}
