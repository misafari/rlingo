package project

import (
	"context"

	"github.com/google/uuid"
	"github.com/misafari/rlingo/internal/project/domain"
)

type Repository interface {
	Create(ctx context.Context, project *domain.Project) (*domain.Project, error)
	Update(ctx context.Context, project *domain.Project) error
	DeleteOneById(ctx context.Context, id uuid.UUID, tenantID uuid.UUID) error
	FetchOneByID(ctx context.Context, projectID, tenantID uuid.UUID) (*domain.Project, error)
	FetchAll(ctx context.Context, tenantID uuid.UUID) ([]*domain.Project, error)
}
