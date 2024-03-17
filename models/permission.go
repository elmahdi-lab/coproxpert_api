package models

//
//import (
//	"github.com/google/uuid"
//)
//
//type Role string
//type EntityType string
//
//const (
//	PropertyEntity     EntityType = "property"
//	BuildingEntity     EntityType = "building"
//	OrganizationEntity EntityType = "organization"
//)
//
//const (
//	AdminRole   Role = "admin_role"
//	ManagerRole Role = "manager_role"
//	UserRole    Role = "user_role"
//)
//
//type Permission struct {
//	ID         uuid.UUID `json:"id" gorm:"type:uuid;primaryKey"`
//	UserID     uuid.UUID
//	User       *User      `json:"user" gorm:"foreignKey:UserID;references:ID;constraint:OnDelete:CASCADE"`
//	Role       Role       `json:"role" gorm:"not null"`
//	EntityType EntityType `json:"entity_type" gorm:"not null"`
//	EntityID   uuid.UUID  `json:"entity_id" gorm:"type:uuid"`
//	BaseModel
//}
//
//// (userID, 'manager_role', 'organization', 'organization_id')
//// (yassine, manager, via_syndic, via_syndic_id)
//// (tajeddine, manager, AUTO, AUTO_ID)
