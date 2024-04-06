package models

import (
	"github.com/google/uuid"
	"ithumans.com/coproxpert/helpers"
	"time"
)

const LockDurationMinutes = 5
const TokenDurationMinutes = 180
const PasswordResetTokenDurationMinutes = 30

type User struct {
	ID uuid.UUID `json:"id" gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`

	Username *string `json:"username" gorm:"unique"`

	FirstName *string `json:"first_name"`
	LastName  *string `json:"last_name"`

	PhoneNumber *string `json:"phone_number"`
	Address     *string `json:"address"`
	City        *string `json:"city"`
	Province    *string `json:"province"`
	ZipCode     *string `json:"zip_code"`
	Country     *string `json:"country"`

	Tries               *int       `json:"tries" gorm:"default:0"`
	LockExpiresAt       *time.Time `json:"lock_expires_at"`
	IsVerified          *bool      `json:"is_verified" gorm:"default:false"`
	PasswordResetToken  *uuid.UUID `json:"password_reset_token"`
	ResetTokenExpiresAt *time.Time `json:"reset_token_expires_at"`

	Password       *string    `json:"password"`
	Token          *uuid.UUID `json:"token"`
	TokenExpiresAt *time.Time `json:"token_expires_at"`

	//Files       *[]File       `json:"files" gorm:"foreignKey:UserID;references:ID;constraint:OnDelete:CASCADE"`
	//Resolutions *[]Resolution `json:"resolutions" gorm:"foreignKey:UserID;references:ID;constraint:OnDelete:CASCADE"`
	//Votes       *[]Vote       `json:"votes" gorm:"foreignKey:UserID;references:ID;constraint:OnDelete:CASCADE"`
	//Complaints  *[]Complaint  `json:"complaints" gorm:"foreignKey:UserID;references:ID;constraint:OnDelete:CASCADE"`

	Permissions []Permission `json:"permissions" gorm:"foreignKey:UserID;references:ID;constraint:OnDelete:CASCADE;preload:true"`

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
	IsLockTimeExpired := u.LockExpiresAt != nil && time.Now().After(*u.LockExpiresAt)
	if u.Tries != nil && *u.Tries >= 5 && !IsLockTimeExpired {
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

func (u *User) Anonymize() {
	u.Password = helpers.StringPointer("***hidden***")
}

func (u *User) isAdmin() bool {
	for _, permission := range u.Permissions {
		if permission.Role == AdminRole {
			return true
		}
	}
	return false
}
