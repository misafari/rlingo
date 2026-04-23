package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	db "github.com/misafari/rlingo/internal/db/generated"
	"github.com/misafari/rlingo/internal/identity"
)

type tenantRepositoryPostgresImpl struct {
	queries *db.Queries
	pool    *pgxpool.Pool
}

func (u *tenantRepositoryPostgresImpl) GetTenantIDByUserID(ctx context.Context, userID uuid.UUID) (uuid.UUID, error) {
	return u.queries.GetTenantIDByUserID(ctx, userID)
}

func NewTenantRepositoryPostgresImpl(queries *db.Queries, pool *pgxpool.Pool) identity.TenantRepository {
	if queries == nil {
		panic("queries is required")
	}

	if pool == nil {
		panic("pool is required")
	}

	return &tenantRepositoryPostgresImpl{
		queries: queries,
		pool:    pool,
	}
}
