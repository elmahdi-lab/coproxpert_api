package models

// Collection (money) think about this
type Collection struct {
	ID     string `json:"id" gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	Amount int    `json:"amount" gorm:"type:integer;not null"`

	Description string `json:"description" gorm:"type:text;not null"`
	UnitGroupID string `json:"unit_group_id" gorm:"type:uuid"`
	UnitID      string `json:"unit_id" gorm:"type:uuid"`

	BaseModel
}
