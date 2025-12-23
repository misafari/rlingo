package translation

import "context"

type TranslationRepository interface {
	CreateNewTranslation(ctx context.Context, translation *Translation) error
}
