package utils

import (
	"golang.org/x/crypto/bcrypt"
)

// HashPassword hashes a plain-text password with MD5.
// (For production use bcrypt or argon2.)
func HashPassword(password string) string {
	b, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(b)
}

func CheckPassword(plain, hashed string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(plain))
	return err == nil
}
