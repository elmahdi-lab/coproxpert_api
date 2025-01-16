package models

import (
	"time"

	"github.com/google/uuid"
)

// SubscriptionType defines the type of subscription.
type SubscriptionType string

type SubscriptionLimitType string

const (
	OrganizationLimit SubscriptionLimitType = "organization"
	UnitLimit         SubscriptionLimitType = "unit"
	UnitGroupLimit    SubscriptionLimitType = "unit_group"
)

const (
	Free       SubscriptionType = "Free"
	Tier1      SubscriptionType = "Tier1"
	Tier2      SubscriptionType = "Tier2"
	Tier3      SubscriptionType = "Tier3"
	Enterprise SubscriptionType = "Enterprise"
)

// SubscriptionTier defines the limits for each subscription type.
type SubscriptionTier struct {
	UnitGroupsLimit int64
	UnitsLimit      int64
}

var SubscriptionTiers = map[SubscriptionType]SubscriptionTier{
	Free: {
		UnitGroupsLimit: 1,
		UnitsLimit:      12,
	},

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
	OrganizationID   *uuid.UUID       `json:"organizationID" gorm:"type:uuid;constraint:OnDelete:CASCADE;"`
	UserID           *uuid.UUID       `json:"userID" gorm:"type:uuid;constraint:OnDelete:CASCADE;"`
	SubscriptionType SubscriptionType `json:"subscriptionType" gorm:"not null, default:'Tier1'"`
	ExpiresAt        *time.Time       `json:"expiresAt"`
	Features         []Feature        `json:"features" gorm:"many2many:subscription_features;"`
	BaseModel
}
