package models

import "github.com/google/uuid"

// Contract holds a contract between the building and a service provider (e.g. cleaning, auth, etc.)
type Contract struct {
	ID uuid.UUID `json:"id" gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	BaseModel
}
