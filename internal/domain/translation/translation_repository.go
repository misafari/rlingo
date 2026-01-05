package translation

import (
	"context"

	"github.com/google/uuid"
)

type Repository interface {
	Create(context.Context, *Translation) error
	FetchAll(context.Context) ([]*Translation, error)
	DeleteOneById(context.Context, uuid.UUID) error
	Update(context.Context, *Translation) error
}
