package userService

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"note_pad/utils"
	"strings"
	"time"
)

// RefreshToken implements [UserService].
func (s *userService) RefreshToken(token string) (string, error) {

	parts := strings.Split(token, ".")
	if len(parts) != 3 {
		return "", fmt.Errorf("invalid token")
	}

	msg := parts[0] + "." + parts[1]
	mac := hmac.New(sha256.New, []byte(s.jwtSecret))
	mac.Write([]byte(msg))
	expected := base64.RawURLEncoding.EncodeToString(mac.Sum(nil))

	if expected != parts[2] {
		return "", fmt.Errorf("invalid signature")
	}

	raw, err := base64.RawURLEncoding.DecodeString(parts[1])
	if err != nil {
		return "", fmt.Errorf("invalid payload")
	}

	var p utils.JWTPayload
	if err := json.Unmarshal(raw, &p); err != nil {
		return "", fmt.Errorf("malformed payload")
	}

	u, err := s.repo.FindByID(p.Sub)

	if err != nil {
		return "", fmt.Errorf("User does not exist")
	}

	if p.Type != utils.RefreshToken {
		return "", fmt.Errorf("this endpoint requires an refresh token")
	}

	// Check token expiration
	if p.Exp > 0 && time.Now().Unix() > p.Exp {
		return "", fmt.Errorf("token expired")
	}

	return utils.GenerateJWT(u, utils.AccessToken, s.jwtSecret, s.jwtExpiryDays)

}
