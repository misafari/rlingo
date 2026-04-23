package identity

import (
	"errors"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

const secret = "4R9tX/yH8Pz3KqL7mN2bV5jS1wE9gA6uX9zR0fC2kL8="

type TokenService interface {
	IssueToken(userID uuid.UUID) (string, time.Time, error)
	ParseToken(raw string) (*Claims, error)
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
	signed, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", time.Time{}, fmt.Errorf("sign token: %w", err)
	}

	return signed, expiresAt, nil
}

func (t *tokenServiceImpl) ParseToken(raw string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(
		raw,
		&Claims{},
		func(t *jwt.Token) (any, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fiber.NewError(fiber.StatusUnauthorized,
					"unexpected signing method: "+t.Header["alg"].(string))
			}
			return []byte(secret), nil
		},
		jwt.WithExpirationRequired(),
		jwt.WithIssuedAt(),
	)
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token claims")
	}
	if claims.UserID == uuid.Nil {
		return nil, errors.New("token missing user id")
	}

	return claims, nil
}

func NewTokenService() TokenService {
	return &tokenServiceImpl{}
}
