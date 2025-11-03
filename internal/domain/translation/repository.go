package translation

import "context"

type Repository interface {
	Create(ctx context.Context, t *Translation) error
	GetByID(ctx context.Context, tenantID, id string) (*Translation, error)
	Update(ctx context.Context, t *Translation) error
	Delete(ctx context.Context, tenantID, id string) error
	ListByProject(ctx context.Context, tenantID, projectID string) ([]*Translation, error)
}
