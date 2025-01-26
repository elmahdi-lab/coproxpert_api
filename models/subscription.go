package models

import (
	"time"

	"github.com/google/uuid"
	"ithumans.com/coproxpert/helpers"
)

// SubscriptionType defines the type of subscription.
type SubscriptionType string

type SubscriptionLimitType string

const (
	UnitLimit      SubscriptionLimitType = "unit"
	UnitGroupLimit SubscriptionLimitType = "unit_group"
)

const (
	Tier1      SubscriptionType = "t1"
	Tier2      SubscriptionType = "t2"
	Tier3      SubscriptionType = "t3"
	Enterprise SubscriptionType = "e"
)

type SubscriptionTier struct {
	UnitGroupsLimit int64
	UnitsLimit      int64
}

var SubscriptionTiers = map[SubscriptionType]SubscriptionTier{
	Tier1: {
		UnitGroupsLimit: 2,
		UnitsLimit:      20,
	},
	Tier2: {
		UnitGroupsLimit: 5,
		UnitsLimit:      50,
	},
	Tier3: {
		UnitGroupsLimit: 15,
		UnitsLimit:      250,
	},
	Enterprise: {
		UnitGroupsLimit: 100,
		UnitsLimit:      10000,
	},
}

type Feature struct {
	ID   uuid.UUID `json:"id" gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	Name string    `json:"name"`
	BaseModel
}

type Subscription struct {
	ID               uuid.UUID        `json:"id" gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	UserID           *uuid.UUID       `json:"userID" gorm:"type:uuid;constraint:OnDelete:CASCADE;"`
	SubscriptionType SubscriptionType `json:"subscriptionType" gorm:"not null, default:'Tier1'"`
	ExpiresAt        *time.Time       `json:"expiresAt"`
	Features         []Feature        `json:"features" gorm:"many2many:subscription_features;"`
	BaseModel
}

func (s *Subscription) CreateTrialSubscription(user *User, subscriptionType SubscriptionType) {
	s.UserID = helpers.UuidPointer(user.ID)
	s.SubscriptionType = subscriptionType
	s.ExpiresAt = helpers.TimePointer(time.Now().AddDate(0, 0, 30))
}

func (s *Subscription) IsExpired() bool {

	if s.ExpiresAt == nil {
		return false
	}

	return s.ExpiresAt.Before(time.Now())
}
