package translation

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/misafari/rlingo/internal/domain/translation"
	"github.com/valyala/fasthttp"
)

type CudTranslationUseCase struct {
	repo translation.Repository
}

func (u *CudTranslationUseCase) Create(ctx context.Context, tr *translation.Translation) error {
	if err := translation.Validate(tr, false); err != nil {
		return err
	}

	return u.repo.Create(ctx, tr)
}

func (u *CudTranslationUseCase) DeleteOneById(ctx *fasthttp.RequestCtx, id uuid.UUID) error {
	if id == uuid.Nil {
		return errors.New("translation id cannot be empty")
	}

	return u.repo.DeleteOneById(ctx, id)
}

func (u *CudTranslationUseCase) Update(ctx context.Context, tr *translation.Translation) error {
	if err := translation.Validate(tr, true); err != nil {
		return err
	}

	return u.repo.Update(ctx, tr)
}

func NewCudTranslationUseCase(repo translation.Repository) *CudTranslationUseCase {
	return &CudTranslationUseCase{
		repo: repo,
	}
}
