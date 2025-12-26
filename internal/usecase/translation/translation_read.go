package translation

import (
	"context"
	"fmt"

	"github.com/misafari/rlingo/internal/domain/translation"
)

type ReadTranslationUseCase struct {
	repo translation.Repository
}

func NewReadTranslationUseCase(repo translation.Repository) *ReadTranslationUseCase {
	return &ReadTranslationUseCase{
		repo: repo,
	}
}

func (u *ReadTranslationUseCase) FetchAll(ctx context.Context) ([]*translation.Translation, error) {
	translations, err := u.repo.FetchAll(ctx)
	if err != nil {
		return nil, fmt.Errorf("usecase CreateNewTranslation: %w", err)
	}

	return translations, nil
}
