package models

import (
	"github.com/google/uuid"
	"time"
)

// SubscriptionType defines the type of subscription.
type SubscriptionType string

const (
	Paid SubscriptionType = "paid"
	Free SubscriptionType = "free"
)

// Subscription represents a subscription entity.
type Subscription struct {
	ID               uuid.UUID         `json:"id" gorm:"primaryKey"`
	OrganizationID   uuid.UUID         `json:"organizationID"`
	Organization     *Organization     `json:"organization" gorm:"foreignKey:OrganizationID;references:ID;constraint:OnDelete:CASCADE"`
	SubscriptionType *SubscriptionType `json:"subscriptionType" gorm:"not null;check:subscription_type IN ('paid', 'free')"`
	ExpiresAt        *time.Time        `json:"expiresAt"`
}
