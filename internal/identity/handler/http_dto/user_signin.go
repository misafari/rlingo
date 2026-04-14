package http_dto

import (
	"time"

	"github.com/misafari/rlingo/internal/identity"
)

type UserSigninRequest struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type UserSignInResponse struct {
	Email       string    `json:"email"`
	FullName    string    `json:"full_name"`
	AccessToken string    `json:"access_token"`
	ExpiresAt   time.Time `json:"expires_at"`
}

func NewUserSignInResponseFromEntity(entity *identity.SignInResult) *UserSignupResponse {
	return &UserSignupResponse{
		Email:       entity.Email,
		FullName:    entity.FullName,
		AccessToken: entity.AccessToken,
		ExpiresAt:   entity.ExpiresAt,
	}
}
