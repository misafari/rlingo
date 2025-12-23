package translation

import (
	"context"
	"errors"
	"fmt"

	"github.com/misafari/rlingo/internal/domain/translation"
)

type TranslationModifyingUseCase struct {
	repo translation.TranslationRepository
}

func NewModifyingUseCase(repo translation.TranslationRepository) *TranslationModifyingUseCase {
	return &TranslationModifyingUseCase{
		repo: repo,
	}
}

func (u *TranslationModifyingUseCase) Create(ctx context.Context, tr *translation.Translation) error {
	if tr.Key == "" {
		return errors.New("translation key cannot be empty")
	}
	if tr.Locale == "" {
		return errors.New("locale must be specified")
	}

	err := u.repo.CreateNewTranslation(ctx, tr)
	if err != nil {
		return fmt.Errorf("usecase CreateNewTranslation: %w", err)
	}

	return nil
}
