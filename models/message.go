package models

import "github.com/google/uuid"

type Message struct {
	ID uuid.UUID `json:"id" gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`

	Content string    `json:"content"`
	UserID  uuid.UUID `json:"user_id" gorm:"type:uuid"`

	IsRead bool `json:"is_read"`

	// ROOM

	BaseModel
}
