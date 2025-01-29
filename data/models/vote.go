package models

import "github.com/google/uuid"

type Vote struct {
	ID uuid.UUID `json:"id" gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`

	UserID       uuid.UUID `json:"user_id" gorm:"type:uuid"`
	ResolutionID uuid.UUID `json:"resolution_id" gorm:"type:uuid"`
	IsApproved   bool      `json:"is_approved" gorm:"default:false"`

	BaseModel
}
