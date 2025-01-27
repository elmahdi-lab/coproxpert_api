package models

import (
	"time"

	"github.com/google/uuid"
	"ithumans.com/coproxpert/helpers"
	"ithumans.com/coproxpert/types"
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

	IsClaimed       *bool `json:"is_claimed" gorm:"default:false"` // when a new user sets a password
	IsEmailVerified *bool `json:"is_email_verified" gorm:"default:false"`
	IsPhoneVerified *bool `json:"is_phone_verified" gorm:"default:false"`

	Tries         *int       `json:"tries" gorm:"default:0"`
	LockExpiresAt *time.Time `json:"lock_expires_at"`

	PasswordResetToken  *uuid.UUID `json:"password_reset_token"`
	ResetTokenExpiresAt *time.Time `json:"reset_token_expires_at"`

	Password *string `json:"password"`
	Token    *string `json:"token" gorm:"-"`

	SignInProvider *types.SignInProvider `json:"sign_in_provider"  gorm:"default:'email'"`
	ProviderID     *string               `json:"provider_id"`

	//Files       *[]File       `json:"files" gorm:"foreignKey:UserID;references:ID;constraint:OnDelete:CASCADE"`
	//Resolutions *[]Resolution `json:"resolutions" gorm:"foreignKey:UserID;references:ID;constraint:OnDelete:CASCADE"`
	//Votes       *[]Vote       `json:"votes" gorm:"foreignKey:UserID;references:ID;constraint:OnDelete:CASCADE"`
	//Complaints  *[]Complaint  `json:"complaints" gorm:"foreignKey:UserID;references:ID;constraint:OnDelete:CASCADE"`

	Permissions  []Permission  `json:"permissions" gorm:"foreignKey:UserID;references:ID;constraint:OnDelete:CASCADE;preload:true"`
	Subscription *Subscription `json:"subscription" gorm:"foreignKey:UserID;references:ID;constraint:OnDelete:CASCADE"`
	BaseModel
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
	u.Tries = helpers.IntPointer(0)
	u.LockExpiresAt = nil
}

func (u *User) IncrementTries() {
	if u.Tries == nil {
		u.Tries = helpers.IntPointer(1)
	} else {
		*u.Tries++
	}
}

func (u *User) Anonymize() {
	u.Password = helpers.StringPointer("***hidden***")
}

func (u *User) IsSuperAdmin() bool {
	for _, permission := range u.Permissions {
		if permission.Role == SuperAdminRole {
			return true
		}
	}
	return false
}

func (u *User) IsPasswordTokenExpired() bool {
	return u.ResetTokenExpiresAt != nil && time.Now().After(*u.ResetTokenExpiresAt)
}
