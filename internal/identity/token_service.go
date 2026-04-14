package identity

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type TokenService interface {
	IssueToken(userID uuid.UUID) (string, time.Time, error)
}

type tokenServiceImpl struct {
}

type Claims struct {
	jwt.RegisteredClaims
	UserID uuid.UUID `json:"uid"`
}

func (t *tokenServiceImpl) IssueToken(userID uuid.UUID) (string, time.Time, error) {
	expiresAt := time.Now().UTC().Add(24 * time.Hour)

	claims := Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   userID.String(),
			IssuedAt:  jwt.NewNumericDate(time.Now().UTC()),
			ExpiresAt: jwt.NewNumericDate(expiresAt),
			Issuer:    "tms",
		},
		UserID: userID,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signed, err := token.SignedString([]byte("4R9tX/yH8Pz3KqL7mN2bV5jS1wE9gA6uX9zR0fC2kL8="))
	if err != nil {
		return "", time.Time{}, fmt.Errorf("sign token: %w", err)
	}

	return signed, expiresAt, nil
}

func NewTokenService() TokenService {
	return &tokenServiceImpl{}
}
