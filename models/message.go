package models

import "github.com/google/uuid"

type Message struct {
	ID      uuid.UUID `json:"id" gorm:"type:uuid; primaryKey"`
	Content string    `json:"content"`
	UserID  uuid.UUID `json:"user_id" gorm:"type:uuid"`
	User    User      `json:"user" gorm:"foreignKey:UserID;references:ID"`
}
