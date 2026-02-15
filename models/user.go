package models

import "time"

// User represents the users table in the database.
// This is the Model layer — holds data structure & domain rules.
type User struct {
	ID        string    `json:"id" gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Name      string    `json:"name" gorm:"type:varchar(255);not null"`
	Email     string    `json:"email" gorm:"type:varchar(255);uniqueIndex;not null"`
	Password  string    `json:"-" gorm:"type:varchar(255);not null"`
	IsOwner   bool      `json:"is_owner" gorm:"default:false"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

type PandingUser struct {
	ID        string    `json:"id" gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Name      string    `json:"name" gorm:"type:varchar(255);not null"`
	Email     string    `json:"email" gorm:"type:varchar(255);not null"`
	Password  string    `json:"-" gorm:"type:varchar(255);not null"`
	Otp       string    `json:"-" gorm:"type:varchar(255);not null"`
	IsOwner   bool      `json:"is_owner" gorm:"default:false"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

type RegisterResponce struct {
	UID string `json:"uid" binding:"required"`
}

type OtpVerifyRequest struct {
	Otp string `json:"otp"     binding:"required"`
	Uid string `json:"uid"     binding:"required"`
}

// ---- Request / Response DTOs ----

type CreateUserRequest struct {
	Name     string `json:"name"     binding:"required"`
	Email    string `json:"email"    binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
	IsOwner  bool   `json:"is_owner"`
}

type LoginRequest struct {
	Email    string `json:"email"    binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type TokenResponse struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refreshToken,omitempty"`
}
