package translation

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/misafari/rlingo/internal/domain/translation"
	"github.com/valyala/fasthttp"
)

type CudTranslationUseCase struct {
	repo translation.Repository
}

func NewCudTranslationUseCase(repo translation.Repository) *CudTranslationUseCase {
	return &CudTranslationUseCase{
		repo: repo,
	}
}

func (u *CudTranslationUseCase) Create(ctx context.Context, tr *translation.Translation) error {
	if tr.Key == "" {
		return errors.New("translation key cannot be empty")
	}
	if tr.Locale == "" {
		return errors.New("locale must be specified")
	}

	err := u.repo.Create(ctx, tr)
	if err != nil {
		return fmt.Errorf("usecase CreateNewTranslation: %w", err)
	}

	return nil
}

func (u *CudTranslationUseCase) DeleteOneById(ctx *fasthttp.RequestCtx, id uuid.UUID) error {
	if id == uuid.Nil {
		return errors.New("translation id cannot be empty")
	}

	return u.repo.DeleteOneById(ctx, id)
}

func (u *CudTranslationUseCase) Update(ctx context.Context, tr *translation.Translation) error {
	if tr.ID == uuid.Nil {
		return errors.New("translation id cannot be empty")
	}
	if tr.Key == "" {
		return errors.New("translation key cannot be empty")
	}
	if tr.Locale == "" {
		return errors.New("locale must be specified")
	}

	return u.repo.Update(ctx, tr)
}
