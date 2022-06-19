package users

import (
	"golang.org/x/crypto/bcrypt"
)

func HashAndSalt(plain string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(plain), bcrypt.MinCost)
	if err != nil {
		return ""
	}

	return string(hash)
}
