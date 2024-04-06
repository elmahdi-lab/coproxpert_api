package models

// Collection (money) think about this
type Collection struct {
	ID          string `json:"id" gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	Amount      int    `json:"amount" gorm:"type:integer;not null"`
	Description string `json:"description" gorm:"type:text;not null"`
	// BuildingID  uuid.UUID `json:"building_id" gorm:"type:uuid;not null"`
}
