package models

import (
	"time"

	"github.com/google/uuid"
	"ithumans.com/coproxpert/data/types"
	"ithumans.com/coproxpert/internals/helpers"
)

const LockDurationMinutes = 5

// TODO: use these values for jwt token duration
const TokenDurationMinutes = 60
const RefreshTokenDurationMinutes = 60 * 24 * 7
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

	Password              *string    `json:"password"`
	RefreshToken          uuid.UUID  `json:"refresh_token" gorm:"type:uuid;default:uuid_generate_v4()"`
	RefreshTokenExpiresAt *time.Time `json:"refresh_token_expires_at"` // TODO the refresh token must have an expiration date, when user created.
	Token                 *string    `json:"token" gorm:"-"`

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

func (u *User) IsRefreshTokenExpired() bool {
	return u.RefreshTokenExpiresAt != nil && time.Now().After(*u.RefreshTokenExpiresAt)
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

func (u *User) IsPasswordTokenExpired() bool {
	return u.ResetTokenExpiresAt != nil && time.Now().After(*u.ResetTokenExpiresAt)
}
