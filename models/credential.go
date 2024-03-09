package models

import (
	"github.com/google/uuid"
	"time"
)

const LockDurationMinutes = 5

type Credential struct {
	ID            uuid.UUID `json:"id" gorm:"type:uuid;primaryKey"`
	UserID        uuid.UUID
	User          *User      `json:"user" gorm:"foreignKey:UserID;references:ID;constraint:OnDelete:CASCADE"`
	Password      *string    `json:"password"`
	Tries         *int       `json:"tries" gorm:"default:0"`
	LockExpiresAt *time.Time `json:"lock_expires_at"`
	BaseModel
}

func (c *Credential) IsLocked() bool {
	if c.Tries != nil && *c.Tries >= 5 {
		return true
	}
	return false
}

func (c *Credential) Lock() {
	if c.LockExpiresAt == nil {
		now := time.Now()
		lockExpiresAt := now.Add(time.Duration(LockDurationMinutes) * time.Minute)
		c.LockExpiresAt = &lockExpiresAt
	}
}

func (c *Credential) Unlock() {
	c.Tries = nil
	c.LockExpiresAt = nil
}

func (c *Credential) IncrementTries() {
	if c.Tries == nil {
		tries := 1
		c.Tries = &tries
	} else {
		*c.Tries++
	}
}
