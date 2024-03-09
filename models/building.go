package models

import "github.com/google/uuid"

type Building struct {
	ID uuid.UUID `json:"id" gorm:"primaryKey"`
}
