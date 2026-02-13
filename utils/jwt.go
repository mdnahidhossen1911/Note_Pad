package utils

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"note_pad/models"
	"strings"
	"time"
)

type JWTPayload struct {
	Sub       string `json:"sub"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	IsOwner   bool   `json:"is_owner"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	Iat       int64  `json:"iat"`
	Exp       int64  `json:"exp"`
}

func GenerateJWT(user *models.User, secret string, expiryDays int) (string, error) {
	header := map[string]string{"alg": "HS256", "typ": "JWT"}
	hJSON, _ := json.Marshal(header)

	now := time.Now()
	payload := JWTPayload{
		Sub:       user.ID,
		Name:      user.Name,
		Email:     user.Email,
		IsOwner:   user.IsOwner,
		CreatedAt: user.CreatedAt.String(),
		UpdatedAt: user.UpdatedAt.String(),
		Iat:       now.Unix(),
		Exp:       now.AddDate(0, 0, expiryDays).Unix(),
	}
	pJSON, err := json.Marshal(payload)
	if err != nil {
		return "", err
	}

	hEnc := base64.RawURLEncoding.EncodeToString(hJSON)
	pEnc := base64.RawURLEncoding.EncodeToString(pJSON)
	msg := hEnc + "." + pEnc

	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write([]byte(msg))
	sig := base64.RawURLEncoding.EncodeToString(mac.Sum(nil))

	return msg + "." + sig, nil
}

func VerifyJWT(token, secret string) (*JWTPayload, error) {
	parts := strings.Split(token, ".")
	if len(parts) != 3 {
		return nil, fmt.Errorf("invalid token")
	}

	msg := parts[0] + "." + parts[1]
	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write([]byte(msg))
	expected := base64.RawURLEncoding.EncodeToString(mac.Sum(nil))

	if expected != parts[2] {
		return nil, fmt.Errorf("invalid signature")
	}

	raw, err := base64.RawURLEncoding.DecodeString(parts[1])
	if err != nil {
		return nil, fmt.Errorf("invalid payload")
	}

	var p JWTPayload
	if err := json.Unmarshal(raw, &p); err != nil {
		return nil, fmt.Errorf("malformed payload")
	}

	// Check token expiration
	if p.Exp > 0 && time.Now().Unix() > p.Exp {
		return nil, fmt.Errorf("token expired")
	}

	return &p, nil
}
