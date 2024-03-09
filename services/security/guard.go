package security

import (
	"github.com/google/uuid"
	"ithumans.com/coproxpert/models"
)

func IsAllowed(u *models.User, a models.AccessLevel, t models.EntityType, e uuid.UUID) bool {

	for _, p := range *u.Permissions {
		if *p.Entity == t && *p.AccessLevel == a {
			if *p.EntityID == e {
				return true
			}
			return false
		}
	}

	return false
}
