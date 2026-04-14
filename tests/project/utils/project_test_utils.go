package utils

import (
	"time"

	"github.com/google/uuid"
	"github.com/misafari/rlingo/internal/project/domain"
)

type ProjectOption func(*domain.Project)

func GenerateProject(opts ...ProjectOption) *domain.Project {
	p := &domain.Project{
		ID:          uuid.New(),
		TenantID:    uuid.New(),
		Name:        "Sample Project",
		Description: "This is a sample project description",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		CreatedBy:   uuid.New(),
		UpdatedBy:   uuid.New(),
	}

	for _, opt := range opts {
		opt(p)
	}
	return p
}

func WithID(id uuid.UUID) ProjectOption {
	return func(p *domain.Project) { p.ID = id }
}

func WithTenantID(id uuid.UUID) ProjectOption {
	return func(p *domain.Project) { p.TenantID = id }
}

func WithCreatedById(id uuid.UUID) ProjectOption {
	return func(p *domain.Project) { p.CreatedBy = id }
}

func WithName(name string) ProjectOption {
	return func(p *domain.Project) { p.Name = name }
}
