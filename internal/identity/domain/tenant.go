package domain

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type TenantPlan string

const (
	PlanFree       TenantPlan = "FREE"
	PlanPro        TenantPlan = "PRO"
	PlanEnterprise TenantPlan = "ENTERPRISE"
)

type TenantStatus string

const (
	TenantStatusActive    TenantStatus = "ACTIVE"
	TenantStatusSuspended TenantStatus = "SUSPENDED"
	TenantStatusDeleted   TenantStatus = "DELETED"
)

type Tenant struct {
	ID        uuid.UUID
	Slug      string
	Name      string
	Plan      TenantPlan
	Status    TenantStatus
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewTenant(slug, name string) (*Tenant, error) {
	if slug == "" {
		return nil, errors.New("slug must not be empty")
	}
	if name == "" {
		return nil, errors.New("name must not be empty")
	}

	now := time.Now().UTC()
	return &Tenant{
		ID:        uuid.New(),
		Slug:      slug,
		Name:      name,
		Plan:      PlanFree,
		Status:    TenantStatusActive,
		CreatedAt: now,
		UpdatedAt: now,
	}, nil
}
