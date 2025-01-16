package models

import "github.com/google/uuid"

// Assembly is a model for general assemblies
type Assembly struct {
	ID        uuid.UUID `json:"id" gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	Name      *string   `json:"name" gorm:"type:varchar(255);not null"`
	StartDate *string   `json:"start_date" gorm:"type:date;not null"`
	EndDate   *string   `json:"end_date" gorm:"type:date;not null"`
	Active    *bool     `json:"active" gorm:"type:boolean;default:true"`
}
