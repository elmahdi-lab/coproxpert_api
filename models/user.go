package models

import (
	"github.com/google/uuid"
	"time"
)

const LockDurationMinutes = 5
const TokenDurationMinutes = 180
const PasswordResetTokenDurationMinutes = 30

type User struct {
	ID                  uuid.UUID  `json:"id" gorm:"type:uuid; primaryKey"`
	Username            *string    `json:"username" gorm:"unique"`
	Password            *string    `json:"password"`
	Tries               *int       `json:"tries" gorm:"default:0"`
	LockExpiresAt       *time.Time `json:"lock_expires_at"`
	IsVerified          *bool      `json:"is_verified" gorm:"default:false"`
	PasswordResetToken  *uuid.UUID `json:"password_reset_token"`
	ResetTokenExpiresAt *time.Time `json:"reset_token_expires_at"`

	Token          *uuid.UUID `json:"token"`
	TokenExpiresAt *time.Time `json:"token_expires_at"`

	Contact     *Contact     `json:"contact" gorm:"foreignKey:UserID;references:ID;constraint:OnDelete:CASCADE"`
	Documents   []Document   `json:"documents" gorm:"foreignKey:UserID;references:ID;constraint:OnDelete:CASCADE"`
	Images      []Image      `json:"images" gorm:"foreignKey:UserID;references:ID;constraint:OnDelete:CASCADE"`
	Resolutions []Resolution `json:"resolutions" gorm:"foreignKey:UserID;references:ID;constraint:OnDelete:CASCADE"`
	Votes       []Vote       `json:"votes" gorm:"foreignKey:UserID;references:ID;constraint:OnDelete:CASCADE"`

	//Permissions []Permission `json:"permissions" gorm:"foreignKey:UserID;references:ID;constraint:OnDelete:CASCADE"`
	BaseModel
}

func (u *User) IsTokenExpired() bool {
	if u.TokenExpiresAt != nil && time.Now().After(*u.TokenExpiresAt) {
		return true
	}
	return false
}

func (u *User) GenerateToken() {
	token := uuid.New()
	u.Token = &token
	u.ExtendValidity()
}

func (u *User) RefreshToken() {
	u.GenerateToken()
}

func (u *User) ExtendValidity() {
	TokenExpiresAt := time.Now().Add(TokenDurationMinutes * time.Minute)
	u.TokenExpiresAt = &TokenExpiresAt
}

func (u *User) IsLocked() bool {
	if u.Tries != nil && *u.Tries >= 5 {
		return true
	}
	return false
}

func (u *User) Lock() {
	if u.LockExpiresAt == nil {
		now := time.Now()
		lockExpiresAt := now.Add(time.Duration(LockDurationMinutes) * time.Minute)
		u.LockExpiresAt = &lockExpiresAt
	}
}

func (u *User) Unlock() {
	u.Tries = nil
	u.LockExpiresAt = nil
}

func (u *User) IncrementTries() {
	if u.Tries == nil {
		tries := 1
		u.Tries = &tries
	} else {
		*u.Tries++
	}
}
