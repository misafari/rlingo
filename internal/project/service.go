package project

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/misafari/rlingo/internal/project/const"
	"github.com/misafari/rlingo/internal/project/domain"
)

type Service struct {
	repository *Repository
}

func (u *Service) Create(ctx context.Context, project *domain.Project) error {
	if project == nil {
		return _const.ErrProjectIsNil
	}

	now := time.Now().UTC()
	id := uuid.New()
	tenantID := uuid.New()
	userID := uuid.New()

	project.ID = id
	project.TenantID = tenantID
	project.CreatedBy = userID
	project.CreatedAt = now
	project.UpdatedAt = now

	if err := project.ValidateWithoutIDCheck(); err != nil {
		return err
	}

	return u.repository.Create(ctx, project)
}

func (u *Service) DeleteOneById(ctx context.Context, projectID uuid.UUID) error {
	tenantID := uuid.New()

	err := u.repository.DeleteOneById(ctx, projectID, tenantID)

	if err != nil {
		return _const.ErrProjectDeletionFailed
	}

	return nil
}

func (u *Service) FetchOneByID(ctx context.Context, ID uuid.UUID) (*domain.Project, error) {
	tenantID := uuid.New()

	projects, err := u.repository.FetchOneByID(ctx, ID, tenantID)

	if err != nil {
		return nil, _const.ErrProjectFetchingFailed
	}

	return projects, nil
}

func (u *Service) FetchAll(ctx context.Context) ([]*domain.Project, error) {
	tenantID := uuid.New()

	projects, err := u.repository.FetchAll(ctx, tenantID)

	if err != nil {
		return nil, _const.ErrProjectFetchingFailed
	}

	return projects, nil
}

func (u *Service) Update(ctx context.Context, project *domain.Project) error {
	if err := project.Validate(); err != nil {
		return err
	}

	err := u.repository.Update(ctx, project)
	if err != nil {
		return err
	}

	return nil
}

func NewProjectService(repository *Repository) *Service {
	return &Service{
		repository: repository,
	}
}
