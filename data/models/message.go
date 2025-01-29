package models

import "github.com/google/uuid"

type Message struct {
	ID uuid.UUID `json:"id" gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`

	Content  string    `json:"content"`
	FromID   uuid.UUID `json:"from_id" gorm:"type:uuid"`
	FromUser *User     `json:"from_user" gorm:"foreignKey:FromID;references:ID;constraint:OnDelete:CASCADE"`
	ToID     uuid.UUID `json:"to_id" gorm:"type:uuid"`
	ToUser   *User     `json:"to_user" gorm:"foreignKey:ToID;references:ID;constraint:OnDelete:CASCADE"`

	IsRead bool `json:"is_read"`

	// ROOM

	BaseModel
}
