package security

import "ithumans.com/coproxpert/models"

func Anonymize(u *models.User) {
	u.Password = nil
	u.Token = nil
	u.TokenExpiresAt = nil
	u.Tries = nil
	u.LockExpiresAt = nil
}
