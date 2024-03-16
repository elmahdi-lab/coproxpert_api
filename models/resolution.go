package models

import "github.com/google/uuid"

type Resolution struct {
	ID             uuid.UUID     `json:"id" gorm:"type:uuid; primaryKey"`
	UserID         uuid.UUID     `json:"user_id" gorm:"type:uuid"`
	User           *User         `json:"user" gorm:"foreignKey:UserID;references:ID;constraint:OnDelete:CASCADE"`
	Organization   *Organization `json:"organization" gorm:"foreignKey:OrganizationID;references:ID;constraint:OnDelete:CASCADE"`
	OrganizationID uuid.UUID     `json:"organization_id" gorm:"type:uuid"`

	Building   *Building `json:"building" gorm:"foreignKey:BuildingID;references:ID;constraint:OnDelete:CASCADE"`
	BuildingID uuid.UUID `json:"building_id" gorm:"type:uuid"`

	IsClosed   bool `json:"is_closed" gorm:"default:false"`
	IsApproved bool `json:"is_approved" gorm:"default:false"`

	PercentageRequired int `json:"percentage_required" gorm:"default:51"`

	Votes []Vote `json:"votes" gorm:"foreignKey:ResolutionID;references:ID;constraint:OnDelete:CASCADE"`

	BaseModel
}
