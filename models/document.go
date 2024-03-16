package models

import "github.com/google/uuid"

type DocumentType string

const (
	Pdf DocumentType = "pdf"
)

type Document struct {
	ID           uuid.UUID    `json:"id" gorm:"type:uuid; primaryKey"`
	Url          string       `json:"url"`
	UserID       uuid.UUID    `json:"user_id" gorm:"type:uuid"`
	User         User         `json:"user" gorm:"foreignKey:UserID;references:ID"`
	DocumentType DocumentType `json:"document_type" gorm:"not null"`
	BaseModel
}
