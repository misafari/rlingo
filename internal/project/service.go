package project

import (
	"context"

	"github.com/google/uuid"
	"github.com/misafari/rlingo/internal/project/domain"
	"github.com/misafari/rlingo/internal/db/generated"
)

type Service struct {
	repository *Repository
}

func (u *Service) Create(ctx context.Context, project *domain.Project) error {
	if project == nil {
		return error2.ErrProjectIsNil
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
		return error2.ErrProjectDeletionFailed
	}

	return nil
}

func (u *Service) FetchAll(ctx context.Context) ([]*domain.Project, error) {
	tenantID := uuid.New()

	projects, err := u.repository.FetchAll(ctx, tenantID)

	if err != nil {
		return nil, error2.ErrProjectFetchingFailed
	}

	return projects, nil
}

func (u *Service) Update(ctx context.Context, project *domain.Project) error {
	if err := project.Validate(true); err != nil {
		return err
	}

	tenantID := uuid.New()
	now := time.Now().UTC()

	_, err := u.repository.Update(ctx, project)
	if err != nil {
		return err
	}

	return nil
}

func NewProjectService(queries *generated.Queries) *Service {
	return &Service{
		queries: queries,
	}
}
