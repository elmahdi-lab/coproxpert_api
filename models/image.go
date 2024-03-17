package models

import "github.com/google/uuid"

type Image struct {
	ID     uuid.UUID `json:"id" gorm:"type:uuid; primaryKey"`
	Url    string    `json:"url"`
	UserID uuid.UUID `json:"user_id" gorm:"type:uuid"`
	User   User      `json:"user" gorm:"foreignKey:UserID;references:ID"`
	BaseModel
}
