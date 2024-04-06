package models

import "github.com/google/uuid"

type Status int8

const (
	Pending  Status = 0
	Accepted Status = 1
	Rejected Status = 2
	Closed   Status = 3
)

type Resolution struct {
	ID             string     `json:"id" gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	UserID         uuid.UUID  `json:"user_id" gorm:"type:uuid"`
	OrganizationID uuid.UUID  `json:"organization_id" gorm:"type:uuid"`
	UnitGroupID    *uuid.UUID `json:"unit_group_id" gorm:"type:uuid"`

	Status Status `json:"status" gorm:"default:0"`

	PercentageRequired int `json:"percentage_required" gorm:"default:51"`

	Votes []Vote `json:"votes" gorm:"foreignKey:ResolutionID;references:ID;constraint:OnDelete:CASCADE"`

	BaseModel
}

func (r *Resolution) IsClosed() bool {
	return r.Status == Closed
}

func (r *Resolution) IsAccepted() bool {
	return r.Status == Accepted
}

func (r *Resolution) IsRejected() bool {
	return r.Status == Rejected
}

func (r *Resolution) IsPending() bool {
	return r.Status == Pending
}

// IterateVotes Iterate Votes iterates over the votes and returns the number of votes that are approved.
func (r *Resolution) IterateVotes() int {
	approved := 0
	rejected := 0

	for _, vote := range r.Votes {
		if vote.IsApproved {
			approved++
		} else {
			rejected++
		}
	}
	return approved / len(r.Votes) * 100
}
