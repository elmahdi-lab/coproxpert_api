package models

import (
	"github.com/google/uuid"
)

type Organization struct {
	ID        uuid.UUID `json:"id" gorm:"type:uuid;primaryKey"`
	Name      string    `json:"name" gorm:"unique"`
	IsEnabled *bool     `json:"is_enabled"`
	BaseModel
}
