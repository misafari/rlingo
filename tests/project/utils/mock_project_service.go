package utils

import (
	"context"

	"github.com/google/uuid"
	"github.com/misafari/rlingo/internal/project"
	"github.com/misafari/rlingo/internal/project/domain"
)

type mockProjectService struct {
	FetchAllFn      func(ctx context.Context) ([]*domain.Project, error)
	FetchOneByIDFn  func(ctx context.Context, id uuid.UUID) (*domain.Project, error)
	DeleteOneByIdFn func(ctx context.Context, id uuid.UUID) error
	CreateFn        func(ctx context.Context, project *domain.Project) (*domain.Project, error)
	UpdateFn        func(ctx context.Context, project *domain.Project) error
}

func (m *mockProjectService) FetchAll(ctx context.Context) ([]*domain.Project, error) {
	if m.FetchAllFn != nil {
		return m.FetchAllFn(ctx)
	}

	return nil, nil
}

func (m *mockProjectService) FetchOneByID(ctx context.Context, id uuid.UUID) (*domain.Project, error) {
	if m.FetchOneByIDFn != nil {
		return m.FetchOneByIDFn(ctx, id)
	}

	return nil, nil
}

func (m *mockProjectService) Create(ctx context.Context, project *domain.Project) (*domain.Project, error) {
	if m.CreateFn != nil {
		return m.CreateFn(ctx, project)
	}

	return nil, nil
}

func (m *mockProjectService) Update(ctx context.Context, project *domain.Project) error {
	if m.UpdateFn != nil {
		return m.UpdateFn(ctx, project)
	}

	return nil
}

func (m *mockProjectService) DeleteOneById(ctx context.Context, id uuid.UUID) error {
	if m.DeleteOneByIdFn != nil {
		return m.DeleteOneByIdFn(ctx, id)
	}

	return nil
}

func NewMockProjectServiceWithFetchAllFn(returnVal []*domain.Project, expectedErr error) project.Service {
	ms := &mockProjectService{}

	ms.FetchAllFn = func(ctx context.Context) ([]*domain.Project, error) {
		return returnVal, expectedErr
	}

	return ms
}

func NewMockProjectServiceWithFetchOneByIDFn(returnVal *domain.Project, expectedErr error) project.Service {
	ms := &mockProjectService{}

	ms.FetchOneByIDFn = func(ctx context.Context, id uuid.UUID) (*domain.Project, error) {
		return returnVal, expectedErr
	}

	return ms
}

func NewMockProjectServiceWithDeleteOneByIdFn(expectedErr error) project.Service {
	ms := &mockProjectService{}

	ms.DeleteOneByIdFn = func(ctx context.Context, id uuid.UUID) error {
		return expectedErr
	}

	return ms
}

func NewMockProjectServiceWithCreateFn(returnVal *domain.Project, expectedErr error) project.Service {
	ms := &mockProjectService{}

	ms.CreateFn = func(ctx context.Context, project *domain.Project) (*domain.Project, error) {
		return returnVal, expectedErr
	}

	return ms
}
