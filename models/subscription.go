package models

import (
	"github.com/google/uuid"
	"time"
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
	Tier1: {
		UnitGroupsLimit: 1,
		UnitsLimit:      15,
	},
	Tier2: {
		UnitGroupsLimit: 5,
		UnitsLimit:      75,
	},
	Tier3: {
		UnitGroupsLimit: 15,
		UnitsLimit:      250,
	},
	Enterprise: {
		UnitGroupsLimit: -1, // -1 for unlimited
		UnitsLimit:      -1, // -1 for unlimited
	},
}

type Feature struct {
	ID   uuid.UUID `json:"id" gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	Name string    `json:"name"`
	BaseModel
}

type Subscription struct {
	ID               uuid.UUID        `json:"id" gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	OrganizationID   uuid.UUID        `json:"organizationID" gorm:"type:uuid"`
	SubscriptionType SubscriptionType `json:"subscriptionType" gorm:"not null, default:'Tier1'"`
	ExpiresAt        time.Time        `json:"expiresAt"`
	Features         []Feature        `json:"features" gorm:"many2many:subscription_features;"`
	BaseModel
}
