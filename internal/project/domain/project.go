package domain

import (
	"time"

	"github.com/google/uuid"
	"github.com/misafari/rlingo/internal/project/const"
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
		return _const.ErrMissingProjectID
	}

	return p.ValidateWithoutIDCheck()
}

func (p *Project) ValidateWithoutIDCheck() error {
	if p.TenantID == uuid.Nil {
		return _const.ErrMissingTenantID
	}

	if p.Name == "" {
		return _const.ErrInvalidName
	}

	if p.CreatedBy == uuid.Nil {
		return _const.ErrMissingCreatedBy
	}

	return nil
}

func NewProject(tenantID, createdBy uuid.UUID, name, description string) (*Project, error) {
	if name == "" {
		return nil, _const.ErrInvalidName
	}

	if tenantID == uuid.Nil {
		return nil, _const.ErrMissingTenantID
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
		return _const.ErrNothingToUpdate
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
