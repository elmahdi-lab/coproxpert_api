package models

import "github.com/google/uuid"

type Vote struct {
	ID           uuid.UUID   `json:"id" gorm:"type:uuid; primaryKey"`
	UserID       uuid.UUID   `json:"user_id" gorm:"type:uuid"`
	User         *User       `json:"user" gorm:"foreignKey:UserID;references:ID;constraint:OnDelete:CASCADE"`
	ResolutionID uuid.UUID   `json:"resolution_id" gorm:"type:uuid"`
	Resolution   *Resolution `json:"resolution" gorm:"foreignKey:ResolutionID;references:ID;constraint:OnDelete:CASCADE"`
	IsApproved   bool        `json:"is_approved" gorm:"default:false"`
	BaseModel
}
