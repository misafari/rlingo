package translation

import (
	"context"

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
	return u.repo.FetchAll(ctx)
}
