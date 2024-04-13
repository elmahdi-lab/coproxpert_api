package models

// Invoice is a model for invoices issued by service providers or suppliers
type Invoice struct {
	ID             string `json:"id" gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	Amount         int    `json:"amount" gorm:"type:integer;not null"`
	Description    string `json:"description" gorm:"type:text;not null"`
	OrganizationID string `json:"organization_id" gorm:"type:uuid"`
	BaseModel
}
