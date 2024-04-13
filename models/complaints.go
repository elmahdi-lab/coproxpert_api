package models

import "github.com/google/uuid"

type Complaint struct {
	ID          string    `json:"id" gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	Title       string    `json:"title"`
	Description string    `json:"description" gorm:"type:text"`
	UserID      uuid.UUID `json:"user_id" gorm:"type:uuid"`

	Files []File `json:"files" gorm:"foreignKey:UserID;references:ID;constraint:OnDelete:CASCADE"`

	IsResolved bool   `json:"is_resolved"`
	ResolvedAt string `json:"resolved_at"`
	Response   string `json:"response" gorm:"type:text"`
	BaseModel
}
