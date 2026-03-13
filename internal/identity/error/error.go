package error

import "errors"

var (
	ErrInvalidSlug = errors.New("identity: slug must not be empty")
	ErrInvalidName = errors.New("identity: name must not be empty")

	ErrSlugAlreadyTaken  = errors.New("identity: slug is already taken")
	ErrEmailAlreadyTaken = errors.New("identity: email is already registered")

	ErrTenantNotFound = errors.New("identity: tenant not found")
	ErrUserNotFound   = errors.New("identity: user not found")

	ErrInvalidCredentials = errors.New("identity: invalid credentials")
	ErrTokenExpired       = errors.New("identity: token has expired")
)
