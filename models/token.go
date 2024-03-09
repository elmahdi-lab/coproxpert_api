package models

import (
	"github.com/google/uuid"
	"time"
)

const ValidUntilDuration = 180

type Token struct {
	ID         uuid.UUID  `json:"id" gorm:"type:uuid;primaryKey"`
	UserID     uuid.UUID  `json:"user_id"`
	User       *User      `json:"user" gorm:"foreignKey:UserID;references:ID;constraint:OnDelete:CASCADE"`
	Token      *string    `json:"token"`
	ValidUntil *time.Time `json:"valid_until"`
	BaseModel
}

func (t *Token) IsExpired() bool {
	if t.ValidUntil != nil && time.Now().After(*t.ValidUntil) {
		return true
	}
	return false
}

func (t *Token) GenerateToken() {
	token := uuid.New().String()
	t.Token = &token
	validUntil := time.Now().Add(ValidUntilDuration * time.Minute)
	t.ValidUntil = &validUntil
}

func (t *Token) RefreshToken() {
	t.GenerateToken()
}

func (t *Token) ExtendValidity() {
	validUntil := time.Now().Add(ValidUntilDuration * time.Minute)
	t.ValidUntil = &validUntil
}
