package models

import "errors"

var (
	ErrUserNotFound    = errors.New("user not found")
	ErrProductNotFound = errors.New("product not found")
	ErrInvalidPassword = errors.New("invalid credentials")
	ErrEmailExists     = errors.New("email already exists")
)
