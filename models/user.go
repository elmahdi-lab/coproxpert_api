package models

import (
	"github.com/google/uuid"
)

type User struct {
	ID          uuid.UUID     `json:"id" gorm:"type:uuid; primaryKey"`
	Username    *string       `json:"username" gorm:"unique"`
	Contacts    *[]Contact    `json:"contacts" gorm:"foreignKey:UserID;references:ID;constraint:OnDelete:CASCADE"`
	Permissions *[]Permission `json:"permissions" gorm:"foreignKey:UserID;references:ID;constraint:OnDelete:CASCADE"`
	IsVerified  *bool         `json:"is_verified" gorm:"default:false"`
	BaseModel
}
