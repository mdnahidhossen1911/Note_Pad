package utils

import (
	"crypto/md5"
	"encoding/hex"
)

// HashPassword hashes a plain-text password with MD5.
// (For production use bcrypt or argon2.)
func HashPassword(password string) string {
	h := md5.Sum([]byte(password))
	return hex.EncodeToString(h[:])
}
