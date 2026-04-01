package domain

import (
	"time"

	"github.com/google/uuid"
	error2 "github.com/misafari/rlingo/internal/project/error"
)

type ProjectStatus string

const (
	ProjectStatusActive   ProjectStatus = "ACTIVE"
	ProjectStatusArchived ProjectStatus = "ARCHIVED"
)

type Project struct {
	ID          uuid.UUID
	Name        string
	Description string
	Status      ProjectStatus
	TenantID    uuid.UUID
	CreatedBy   uuid.UUID
	CreatedAt   time.Time
	UpdatedBy   uuid.UUID
	UpdatedAt   time.Time
}

func (p *Project) Validate() error {
	if p.ID == uuid.Nil {
		return error2.ErrMissingProjectID
	}

	return p.ValidateWithoutIDCheck()
}

func (p *Project) ValidateWithoutIDCheck() error {
	if p.TenantID == uuid.Nil {
		return error2.ErrMissingTenantID
	}

	if p.Name == "" {
		return error2.ErrInvalidName
	}

	if p.CreatedBy == uuid.Nil {
		return error2.ErrMissingCreatedBy
	}

	return nil
}

func NewProject(tenantID, createdBy uuid.UUID, name, description string) (*Project, error) {
	if name == "" {
		return nil, error2.ErrInvalidName
	}

	if tenantID == uuid.Nil {
		return nil, error2.ErrMissingTenantID
	}

	now := time.Now().UTC()
	return &Project{
		ID:          uuid.New(),
		TenantID:    tenantID,
		Name:        name,
		Description: description,
		Status:      ProjectStatusActive,
		CreatedBy:   createdBy,
		CreatedAt:   now,
		UpdatedAt:   now,
	}, nil
}

func (p *Project) Update(name, description string) error {
	if name == "" && description == "" {
		return error2.ErrNothingToUpdate
	}

	if name != "" {
		p.Name = name
	}

	if description != "" {
		p.Description = description
	}

	p.UpdatedAt = time.Now().UTC()
	return nil
}
