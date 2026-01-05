package project

import (
	"context"

	"github.com/google/uuid"
	"github.com/misafari/rlingo/internal/domain/project"
)

type CrudProjectUseCase struct {
	repo project.Repository
}

func (u *CrudProjectUseCase) Create(ctx context.Context, p *project.Project) error {
	if err := project.Validate(p, false); err != nil {
		return err
	}

	return u.repo.Create(ctx, p)
}

func (u *CrudProjectUseCase) FetchAll(ctx context.Context) ([]*project.Project, error) {
	return u.repo.FetchAll(ctx)
}

func (u *CrudProjectUseCase) DeleteOneById(ctx context.Context, uuid uuid.UUID) error {
	return u.repo.DeleteOneById(ctx, uuid)
}

func (u *CrudProjectUseCase) Update(ctx context.Context, p *project.Project) error {
	if err := project.Validate(p, true); err != nil {
		return err
	}

	return u.repo.Update(ctx, p)
}

func NewCrudProjectUseCase(repo project.Repository) *CrudProjectUseCase {
	return &CrudProjectUseCase{
		repo: repo,
	}
}
