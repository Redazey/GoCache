package errorz

import (
	"errors"
)

var (
	ErrTokenExpired = errors.New("tokenExpiredError")
	ErrUserExists   = errors.New("userAlreadyRegisteredError")
	ErrValidation   = errors.New("validationError")
	ErrNilCacheData = errors.New("nilCacheData")
	ErrUserNotFound = errors.New("userNotFound")
)
