package errors

import "errors"

var (
	ErrEmailRequired      = errors.New("email is required")
	ErrEmailInvalid       = errors.New("email format is invalid")
	ErrEmailTaken         = errors.New("email is already registered")
	ErrFullNameRequired   = errors.New("full name is required")
	ErrFullNameTooLong    = errors.New("full name must be 255 characters or fewer")
	ErrPasswordTooShort   = errors.New("password must be at least 8 characters")
	ErrPasswordTooLong    = errors.New("password must be 72 characters or fewer") // bcrypt limit
	ErrInvalidCredentials = errors.New("email or password is incorrect")
	ErrUserNotFound       = errors.New("user not found")
)
