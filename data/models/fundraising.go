package models

// Fundraising is a model for fundraising campaigns
type Fundraising struct {
	ID          string  `json:"id" gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	Amount      float64 `json:"amount" gorm:"not null"`
	Description string  `json:"description" gorm:"type:text;not null"`
	OwnerID     string  `json:"owner_id" gorm:"type:uuid;not null"`
	UnitGroupID string  `json:"unit_group_id" gorm:"type:uuid;not null"`
	BaseModel
}
