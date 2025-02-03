package models

import (
	"time"

	"github.com/google/uuid"
	"ithumans.com/coproxpert/internals/helpers"
)

// SubscriptionTier defines the type of subscription.
type SubscriptionTier string

// SubscriptionLimitType defines the type of subscription limit.
type SubscriptionLimitType string

const (
	UnitLimit      SubscriptionLimitType = "unit"
	UnitGroupLimit SubscriptionLimitType = "unit_group"
)

const (
	Tier1      SubscriptionTier = "t1"
	Tier2      SubscriptionTier = "t2"
	Tier3      SubscriptionTier = "t3"
	Enterprise SubscriptionTier = "e"
)

// SubscriptionTierLimits defines the limits for each subscription tier.
type SubscriptionTierLimits struct {
	UnitGroupsLimit int64
	UnitsLimit      int64
}

// SubscriptionTierConfigs maps each subscription tier to its limits.
var SubscriptionTierConfigs = map[SubscriptionTier]SubscriptionTierLimits{
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

// Feature represents a feature available in a subscription.
type Feature struct {
	ID   uuid.UUID `json:"id" gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	Name string    `json:"name"`
	BaseModel
}

// Subscription represents a user's subscription.
type Subscription struct {
	ID        uuid.UUID        `json:"id" gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	UserID    *uuid.UUID       `json:"userID" gorm:"type:uuid;constraint:OnDelete:CASCADE;"`
	Tier      SubscriptionTier `json:"tier" gorm:"not null;default:'t1'"`
	ExpiresAt *time.Time       `json:"expiresAt"`
	Features  []Feature        `json:"features" gorm:"many2many:subscription_features;"`
	BaseModel
}

// InitializeTrialSubscription sets up a trial subscription for a user.
func (s *Subscription) InitializeTrialSubscription(user *User, tier SubscriptionTier) {
	s.UserID = helpers.UuidPointer(user.ID)
	s.Tier = tier
	s.ExpiresAt = helpers.TimePointer(time.Now().AddDate(0, 0, 30)) // Trial expires in 30 days
}

// IsExpired checks if the subscription has expired.
func (s *Subscription) IsExpired() bool {
	if s.ExpiresAt == nil {
		return false
	}
	return s.ExpiresAt.Before(time.Now())
}
