package http_dto

import (
	"time"

	"github.com/misafari/rlingo/internal/identity"
)

type UserSignupRequest struct {
	Email    string `json:"email" validate:"required,min=3,max=100"`
	Password string `json:"password" validate:"required,min=16,max=70"`
	FullName string `json:"full_name" validate:"required,min=3,max=100"`
}

type UserSignupResponse struct {
	Email       string    `json:"email"`
	FullName    string    `json:"full_name"`
	AccessToken string    `json:"access_token"`
	ExpiresAt   time.Time `json:"expires_at"`
}

func NewUserSignupResponseFromEntity(entity *identity.SignUpResult) *UserSignupResponse {
	return &UserSignupResponse{
		Email:       entity.Email,
		FullName:    entity.FullName,
		AccessToken: entity.AccessToken,
		ExpiresAt:   entity.ExpiresAt,
	}
}
