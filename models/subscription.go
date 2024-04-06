package models

import (
	"github.com/google/uuid"
	"time"
)

// SubscriptionType defines the type of subscription.
type SubscriptionType string

// SubscriptionTier defines the limits for each subscription type.
type SubscriptionTier struct {
	Name           SubscriptionType
	BuildingsLimit int
	UnitsLimit     int
}

var (
	Free = SubscriptionTier{
		Name:           "Free",
		BuildingsLimit: 1,
		UnitsLimit:     10,
	}
	Tier1 = SubscriptionTier{
		Name:           "Tier1",
		BuildingsLimit: 10,
		UnitsLimit:     100,
	}
	Tier2 = SubscriptionTier{
		Name:           "Tier2",
		BuildingsLimit: 100,
		UnitsLimit:     1000,
	}
	Tier3 = SubscriptionTier{
		Name:           "Tier3",
		BuildingsLimit: 1000,
		UnitsLimit:     10000,
	}
	Enterprise = SubscriptionTier{
		Name:           "Enterprise",
		BuildingsLimit: -1, // -1 for unlimited
		UnitsLimit:     -1, // -1 for unlimited
	}
)

type Feature struct {
	ID uuid.UUID `json:"id" gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`

	Name string `json:"name"`

	BaseModel
}

type Subscription struct {
	ID uuid.UUID `json:"id" gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`

	OrganizationID   uuid.UUID        `json:"organizationID" gorm:"type:uuid"`
	SubscriptionType SubscriptionType `json:"subscriptionType" gorm:"not null, default:'Free'"`
	ExpiresAt        time.Time        `json:"expiresAt"`
	Features         []Feature        `json:"features" gorm:"many2many:subscription_features;"`
	BaseModel
}
