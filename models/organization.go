package models

import (
	"github.com/google/uuid"
	"strconv"
	"time"
)

type Organization struct {
	ID        uuid.UUID `json:"id" gorm:"type:uuid;primaryKey"`
	Name      string    `json:"name" gorm:"unique"`
	IsEnabled *bool     `json:"is_enabled" gorm:"default:true"`
	BaseModel
}

// GenerateName Generates a name with this format: SELF-[azAZ09]{6}-timestamp
func (o *Organization) GenerateName() {
	unix := time.Now().Unix()
	unixStr := strconv.FormatInt(unix, 10)
	random := uuid.New().String()[:6]
	o.Name = "SELF-" + random + "-" + unixStr
}
