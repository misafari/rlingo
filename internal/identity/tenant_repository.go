package identity

import (
	"context"

	"github.com/google/uuid"
)

type TenantRepository interface {
	GetTenantIDByUserID(ctx context.Context, userID uuid.UUID) (uuid.UUID, error)
}
