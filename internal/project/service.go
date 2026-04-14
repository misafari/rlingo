package project

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/misafari/rlingo/internal/project/const"
	"github.com/misafari/rlingo/internal/project/domain"
)

type Service interface {
	FetchAll(c context.Context) ([]*domain.Project, error)
	FetchOneByID(c context.Context, id uuid.UUID) (*domain.Project, error)
	Create(c context.Context, project *domain.Project) (*domain.Project, error)
	Update(c context.Context, project *domain.Project) error
	DeleteOneById(c context.Context, projectID uuid.UUID) error
}

type serviceImpl struct {
	repository Repository
}

func (u *serviceImpl) Create(ctx context.Context, project *domain.Project) (*domain.Project, error) {
	if project == nil {
		return nil, _const.ErrProjectIsNil
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
		return nil, err
	}

	savedProject, err := u.repository.Create(ctx, project)
	if err != nil {
		return nil, err
	}

	return savedProject, nil
}

func (u *serviceImpl) DeleteOneById(ctx context.Context, projectID uuid.UUID) error {
	tenantID := uuid.New()

	err := u.repository.DeleteOneById(ctx, projectID, tenantID)

	if err != nil {
		return _const.ErrProjectDeletionFailed
	}

	return nil
}

func (u *serviceImpl) FetchOneByID(ctx context.Context, ID uuid.UUID) (*domain.Project, error) {
	tenantID := uuid.New()

	projects, err := u.repository.FetchOneByID(ctx, ID, tenantID)

	if err != nil {
		return nil, _const.ErrProjectFetchingFailed
	}

	return projects, nil
}

func (u *serviceImpl) FetchAll(ctx context.Context) ([]*domain.Project, error) {
	tenantID := uuid.New()

	projects, err := u.repository.FetchAll(ctx, tenantID)

	if err != nil {
		return nil, _const.ErrProjectFetchingFailed
	}

	return projects, nil
}

func (u *serviceImpl) Update(ctx context.Context, project *domain.Project) error {
	if err := project.Validate(); err != nil {
		return err
	}

	err := u.repository.Update(ctx, project)
	if err != nil {
		return err
	}

	return nil
}

func NewProjectService(repository Repository) Service {
	return &serviceImpl{
		repository: repository,
	}
}
