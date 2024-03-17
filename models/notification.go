package models

import "github.com/google/uuid"

type Notification struct {
	ID      uuid.UUID `json:"id" gorm:"type:uuid; primaryKey"`
	UserID  uuid.UUID `json:"user_id" gorm:"type:uuid"`
	User    *User     `json:"user" gorm:"foreignKey:UserID;references:ID;constraint:OnDelete:CASCADE"`
	IsRead  bool      `json:"is_read" gorm:"default:false"`
	Message string    `json:"message"`
	BaseModel
}
