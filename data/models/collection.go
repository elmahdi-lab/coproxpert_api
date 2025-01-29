package models

import "github.com/google/uuid"

// Collection (money) think about this
type Collection struct {
	ID string `json:"id" gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`

	Description string `json:"description" gorm:"type:text;not null"`

	OwnerID uuid.UUID `json:"owner_id" gorm:"type:uuid;not null"`
	UnitID  uuid.UUID `json:"unit_id" gorm:"type:uuid;not null"`
	Amount  float64   `json:"amount" gorm:"not null"`
	BaseModel
}
