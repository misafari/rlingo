package domain

import (
	"regexp"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/google/uuid"
	"github.com/misafari/rlingo/internal/identity/errors"
	"golang.org/x/crypto/bcrypt"
)

type UsersStatus string

const (
	UsersStatusACTIVE    UsersStatus = "ACTIVE"
	UsersStatusSUSPENDED UsersStatus = "SUSPENDED"
	UsersStatusDELETED   UsersStatus = "DELETED"
)

type User struct {
	ID           uuid.UUID
	Email        string
	IsSso        bool
	FullName     string
	PasswordHash string
	Status       UsersStatus
	LastLoginAt  time.Time
	CreatedAt    time.Time
}

var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)

func (u *User) ValidateSignUpInput() error {
	email := strings.TrimSpace(u.Email)
	fullName := strings.TrimSpace(u.FullName)

	if email == "" {
		return errors.ErrEmailRequired
	}
	if !emailRegex.MatchString(email) {
		return errors.ErrEmailInvalid
	}
	if fullName == "" {
		return errors.ErrFullNameRequired
	}
	if utf8.RuneCountInString(fullName) > 255 {
		return errors.ErrFullNameTooLong
	}
	if len(u.PasswordHash) < 8 {
		return errors.ErrPasswordTooShort
	}
	if len(u.PasswordHash) > 72 {
		return errors.ErrPasswordTooLong
	}

	return nil
}

func (u *User) HashPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return err
	}
	u.PasswordHash = string(bytes)
	return nil
}

func (u *User) CheckPassword(providedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(providedPassword))
}
