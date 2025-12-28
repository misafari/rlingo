package translation

import (
	"context"

	"github.com/google/uuid"
)

type Repository interface {
	Create(ctx context.Context, translation *Translation) error
	FetchAll(context.Context) ([]*Translation, error)
	DeleteOneById(context.Context, uuid.UUID) error
	Update(context.Context, *Translation) error
}
