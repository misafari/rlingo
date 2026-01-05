package translation_key

import (
	"context"

	"github.com/google/uuid"
)

type Repository interface {
	Create(context.Context, *Key) error
	FetchAll(context.Context) ([]*Key, error)
	DeleteOneById(context.Context, uuid.UUID) error
	Update(context.Context, *Key) error
}
