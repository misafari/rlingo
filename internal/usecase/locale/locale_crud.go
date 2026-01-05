package locale

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/misafari/rlingo/internal/domain/locale"
)

type CrudLocaleUseCase struct {
	repo locale.Repository
}

func (u *CrudLocaleUseCase) Create(ctx context.Context, entity *locale.Locale) error {
	if err := locale.Validate(entity, false); err != nil {
		return err
	}

	return u.repo.Create(ctx, entity)
}

func (u *CrudLocaleUseCase) FetchAll(ctx context.Context) ([]*locale.Locale, error) {
	return u.repo.FetchAll(ctx)
}

func (u *CrudLocaleUseCase) DeleteOneById(ctx context.Context, id uuid.UUID) error {
	if id == uuid.Nil {
		return errors.New("id is nil")
	}

	return u.repo.DeleteOneById(ctx, id)
}

func (u *CrudLocaleUseCase) Update(ctx context.Context, entity *locale.Locale) error {
	if err := locale.Validate(entity, true); err != nil {
		return err
	}

	return u.repo.Update(ctx, entity)
}

func NewCrudLocaleUseCase(repo locale.Repository) *CrudLocaleUseCase {
	return &CrudLocaleUseCase{
		repo: repo,
	}
}
