package models

import "github.com/google/uuid"

type Complaint struct {
	ID          uuid.UUID `json:"id" gorm:"type:uuid; primaryKey"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	IsResolved  bool      `json:"is_resolved" gorm:"default:false"`
	UserID      uuid.UUID `json:"user_id" gorm:"type:uuid"`
	User        User      `json:"user" gorm:"foreignKey:UserID;references:ID"`

	Images    []Image    `json:"images" gorm:"foreignKey:UserID;references:ID;constraint:OnDelete:CASCADE"`
	Documents []Document `json:"documents" gorm:"foreignKey:UserID;references:ID;constraint:OnDelete:CASCADE"`

	BaseModel
}
