package middleware

import (
	"errors"
	"log/slog"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/misafari/rlingo/internal/identity"
	errors2 "github.com/misafari/rlingo/internal/share/errors"
)

type Claims struct {
	jwt.RegisteredClaims
	UserID uuid.UUID `json:"uid"`
}

type AuthMiddleware struct {
	tokenService  identity.TokenService
	tenantService identity.TenantService
	log           *slog.Logger
}

func (m *AuthMiddleware) TokenValidationFilter() fiber.Handler {
	return func(c *fiber.Ctx) error {
		raw := extractBearerToken(c)
		if raw == "" {
			return errors2.Unauthorized("missing or malformed Authorization header")
		}

		claims, err := m.tokenService.ParseToken(raw)
		if err != nil {
			m.log.Debug("auth middleware: invalid token",
				"error", err.Error(),
				"request_id", c.GetRespHeader(fiber.HeaderXRequestID),
				"path", c.Path(),
			)
			return tokenError(c, err)
		}

		SetAuth(c, AuthContext{
			UserID:    claims.UserID,
			ExpiresAt: claims.ExpiresAt.Time,
		})

		return c.Next()
	}
}

func (m *AuthMiddleware) TenantFilter() fiber.Handler {
	return func(c *fiber.Ctx) error {
		auth := MustGetAuth(c)

		tenantID, err := m.tenantService.GetTenantIDByUserID(auth.UserID)
		if err != nil {
			log.Debug("tenant middleware: lookup failed",
				"error", err.Error(),
			)
			return errors2.Unauthorized("tenant not found or access denied")
		}

		auth.TenantID = tenantID
		SetAuth(c, auth)

		return c.Next()
	}
}

func NewAuthMiddleware(
	tokenService identity.TokenService,
	tenantService identity.TenantService,
	log *slog.Logger,
) *AuthMiddleware {
	return &AuthMiddleware{
		tokenService:  tokenService,
		tenantService: tenantService,
		log:           log,
	}
}

type TenantLookupFn func(c *fiber.Ctx, slug string) (uuid.UUID, error)

func extractBearerToken(c *fiber.Ctx) string {
	header := c.Get(fiber.HeaderAuthorization)
	if !strings.HasPrefix(header, "Bearer ") {
		return ""
	}
	token := strings.TrimPrefix(header, "Bearer ")
	token = strings.TrimSpace(token)
	if token == "" {
		return ""
	}
	return token
}

func tokenError(c *fiber.Ctx, err error) error {
	switch {
	case errors.Is(err, jwt.ErrTokenExpired):
		return RespondError(c, fiber.StatusUnauthorized, "token has expired")
	case errors.Is(err, jwt.ErrTokenNotValidYet):
		return RespondError(c, fiber.StatusUnauthorized, "token is not yet valid")
	case errors.Is(err, jwt.ErrTokenMalformed):
		return RespondError(c, fiber.StatusUnauthorized, "malformed token")
	default:
		return RespondError(c, fiber.StatusUnauthorized, "invalid token")
	}
}

func RespondError(c *fiber.Ctx, status int, msg string) error {
	return c.Status(status).JSON(errorResponse{Error: msg})
}

type errorResponse struct {
	Error string `json:"error"`
}
