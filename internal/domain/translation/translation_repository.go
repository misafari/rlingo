package translation

import (
	"context"
)

type TranslationRepository interface {
	Create(ctx context.Context, translation *Translation) error
	FetchAll(context.Context) ([]*Translation, error)
	DeleteOneById(context.Context, string) error
	Update(context.Context, *Translation) error
}
