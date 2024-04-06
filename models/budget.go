package models

import "github.com/google/uuid"

// Budget is a model for budgets, they may be provisional or definitive
type Budget struct {
	ID uuid.UUID `json:"id" gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
}
