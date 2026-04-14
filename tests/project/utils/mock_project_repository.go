package utils

import (
	"context"

	"github.com/google/uuid"
	"github.com/misafari/rlingo/internal/project"
	"github.com/misafari/rlingo/internal/project/domain"
)

type mockRepository struct {
	FetchAllFn      func(ctx context.Context, tenantID uuid.UUID) ([]*domain.Project, error)
	FetchOneByIDFn  func(ctx context.Context, id uuid.UUID) (*domain.Project, error)
	DeleteOneByIdFn func(ctx context.Context, id uuid.UUID, tenantID uuid.UUID) error
	CreateFn        func(ctx context.Context, project *domain.Project) (*domain.Project, error)
	UpdateFn        func(ctx context.Context, project *domain.Project) error
}

func (m *mockRepository) Create(ctx context.Context, project *domain.Project) (*domain.Project, error) {
	if m.CreateFn != nil {
		return m.CreateFn(ctx, project)
	}

	panic("pass the function to mock repository.Create")
}

func (m *mockRepository) Update(ctx context.Context, project *domain.Project) error {
	if m.UpdateFn != nil {
		return m.UpdateFn(ctx, project)
	}

	panic("pass the function to mock repository.Update")
}

func (m *mockRepository) DeleteOneById(ctx context.Context, id uuid.UUID, tenantID uuid.UUID) error {
	if m.DeleteOneByIdFn != nil {
		return m.DeleteOneByIdFn(ctx, id, tenantID)
	}

	panic("pass the function to mock repository.DeleteOneById")
}

func (m *mockRepository) FetchOneByID(ctx context.Context, projectID, tenantID uuid.UUID) (*domain.Project, error) {
	if m.FetchOneByIDFn != nil {
		return m.FetchOneByIDFn(ctx, projectID)
	}

	panic("pass the function to mock repository.FetchOneByID")
}

func (m *mockRepository) FetchAll(ctx context.Context, tenantID uuid.UUID) ([]*domain.Project, error) {
	if m.FetchAllFn != nil {
		return m.FetchAllFn(ctx, tenantID)
	}

	panic("pass the function to mock repository.FetchAllFn")
}

func NewMockProjectRepositoryWithFetchAllFn(returnVal []*domain.Project, expectedErr error) project.Repository {
	ms := &mockRepository{}

	ms.FetchAllFn = func(ctx context.Context, tenantID uuid.UUID) ([]*domain.Project, error) {
		return returnVal, expectedErr
	}

	return ms
}

func NewMockProjectRepositoryWithFetchOneByIDFn(returnVal *domain.Project, expectedErr error) project.Repository {
	ms := &mockRepository{}

	ms.FetchOneByIDFn = func(ctx context.Context, id uuid.UUID) (*domain.Project, error) {
		return returnVal, expectedErr
	}

	return ms
}

func NewMockProjectRepositoryWithDeleteOneByIdFn(expectedErr error) project.Repository {
	ms := &mockRepository{}

	ms.DeleteOneByIdFn = func(ctx context.Context, id uuid.UUID, tenantID uuid.UUID) error {
		return expectedErr
	}

	return ms
}

func NewMockProjectRepositoryWithCreateFn(returnVal *domain.Project, expectedErr error) project.Repository {
	ms := &mockRepository{}

	ms.CreateFn = func(ctx context.Context, project *domain.Project) (*domain.Project, error) {
		return returnVal, expectedErr
	}

	return ms
}
