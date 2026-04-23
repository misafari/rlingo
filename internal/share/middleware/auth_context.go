package middleware

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type AuthContext struct {
	UserID    uuid.UUID
	TenantID  uuid.UUID
	ExpiresAt time.Time
}

const authContextKey = "auth"

func SetAuth(c *fiber.Ctx, auth AuthContext) {
	c.Locals(authContextKey, auth)
}

func GetAuth(c *fiber.Ctx) (AuthContext, bool) {
	val := c.Locals(authContextKey)
	if val == nil {
		return AuthContext{}, false
	}
	auth, ok := val.(AuthContext)
	return auth, ok
}

func MustGetAuth(c *fiber.Ctx) AuthContext {
	auth, ok := GetAuth(c)
	if !ok {
		panic("platform: MustGetAuth called on unauthenticated context")
	}
	return auth
}
