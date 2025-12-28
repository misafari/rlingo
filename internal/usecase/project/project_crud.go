package project

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/misafari/rlingo/internal/domain/project"
)

type CrudProjectUseCase struct {
	repo project.Repository
}

func (u *CrudProjectUseCase) Create(ctx context.Context, tr *project.Project) error {
	if tr.Name == "" {
		return errors.New("project name cannot be empty")
	}

	return u.repo.Create(ctx, tr)
}

func (u *CrudProjectUseCase) FetchAll(ctx context.Context) ([]*project.Project, error) {
	all, err := u.repo.FetchAll(ctx)
	if err != nil {
		return nil, err
	}

	return all, nil
}

func (u *CrudProjectUseCase) DeleteOneById(ctx context.Context, uuid uuid.UUID) error {
	return u.repo.DeleteOneById(ctx, uuid)
}

func (u *CrudProjectUseCase) Update(ctx context.Context, entity *project.Project) error {
	return u.repo.Update(ctx, entity)
}

func NewCrudProjectUseCase(repo project.Repository) *CrudProjectUseCase {
	return &CrudProjectUseCase{
		repo: repo,
	}
}
