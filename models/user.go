package models

import (
	"github.com/google/uuid"
)

type User struct {
	ID       uuid.UUID `json:"id" gorm:"primaryKey"`
	Username *string   `json:"username" gorm:"unique"`
}
