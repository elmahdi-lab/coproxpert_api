package models

import "github.com/google/uuid"

type Complaint struct {
	ID          uuid.UUID `json:"id" gorm:"type:uuid; primaryKey"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	UserID      uuid.UUID `json:"user_id" gorm:"type:uuid"`

	// TODO: Add files
	// Files []File `json:"files" gorm:"foreignKey:UserID;references:ID;constraint:OnDelete:CASCADE"`

	IsResolved bool `json:"is_resolved"`

	//Message []Message `json:"messages" gorm:"foreignKey:UserID;references:ID;constraint:OnDelete:CASCADE"`

	BaseModel
}
