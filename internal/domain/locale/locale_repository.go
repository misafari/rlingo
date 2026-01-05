package locale

import (
	"context"

	"github.com/google/uuid"
)

type Repository interface {
	Create(context.Context, *Locale) error
	FetchAll(context.Context) ([]*Locale, error)
	DeleteOneById(context.Context, uuid.UUID) error
	Update(context.Context, *Locale) error
}
