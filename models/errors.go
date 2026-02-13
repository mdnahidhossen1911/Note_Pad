package models

import "errors"

var (
	ErrInvalidRequest     = errors.New("invalid request")
	ErrMissingField       = errors.New("missing required field")
	ErrInvalidID          = errors.New("invalid id")
	ErrUnauthorized       = errors.New("unauthorized")
	ErrForbidden          = errors.New("forbidden")
	ErrInternalServer     = errors.New("internal server error")
	ErrServiceUnavailable = errors.New("service unavailable")
	ErrTooManyRequests    = errors.New("too many requests")

	ErrUserNotFound       = errors.New("user not found")
	ErrUserAlreadyExists  = errors.New("user already exists")
	ErrEmailExists        = errors.New("email already exists")
	ErrInvalidEmail       = errors.New("invalid email address")
	ErrInvalidPassword    = errors.New("invalid credentials")
	ErrPasswordTooShort   = errors.New("password is too short")
	ErrPasswordTooWeak    = errors.New("password is too weak")
	ErrAccountDisabled    = errors.New("account is disabled")
	ErrAccountNotVerified = errors.New("account is not verified")

	ErrTokenMissing = errors.New("token missing")
	ErrTokenInvalid = errors.New("invalid token")
	ErrTokenExpired = errors.New("token expired")
	ErrTokenRevoked = errors.New("token revoked")

	ErrOTPRequired     = errors.New("otp is required")
	ErrOTPInvalid      = errors.New("invalid otp")
	ErrOTPExpired      = errors.New("otp expired")
	ErrOTPAlreadyUsed  = errors.New("otp already used")
	ErrOTPNotSent      = errors.New("otp not sent")
	ErrOTPLimitReached = errors.New("otp request limit reached")
	ErrOTPCooldown     = errors.New("please wait before requesting a new otp")
	ErrEmailSendFailed = errors.New("failed to send email")

	ErrProductNotFound = errors.New("product not found")

	ErrDBInsertFailed = errors.New("failed to create record")
	ErrDBUpdateFailed = errors.New("failed to update record")
	ErrDBDeleteFailed = errors.New("failed to delete record")
	ErrDBQueryFailed  = errors.New("failed to fetch record")

	ErrInvalidPhone  = errors.New("invalid phone number")
	ErrInvalidName   = errors.New("invalid name")
	ErrInvalidRole   = errors.New("invalid role")
	ErrInvalidStatus = errors.New("invalid status")

	ErrFileRequired    = errors.New("file is required")
	ErrFileTooLarge    = errors.New("file too large")
	ErrInvalidFileType = errors.New("invalid file type")
)
