package identity

import (
	"context"

	"github.com/google/uuid"
)

type TenantService interface {
	GetTenantIDByUserID(ctx context.Context, userID uuid.UUID) (uuid.UUID, error)
}

type tenantServiceImpl struct {
	repository TenantRepository
}

func (t *tenantServiceImpl) GetTenantIDByUserID(ctx context.Context, userID uuid.UUID) (uuid.UUID, error) {
	return t.repository.GetTenantIDByUserID(ctx, userID)
}

func NewTenantService(repository TenantRepository) TenantService {
	return &tenantServiceImpl{
		repository: repository,
	}
}
