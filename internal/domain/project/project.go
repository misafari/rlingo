package project

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type Project struct {
	ID        uuid.UUID
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func Validate(p *Project, idCheck bool) error {
	if p == nil {
		return errors.New("project is nil")
	}

	if idCheck && p.ID == uuid.Nil {
		return errors.New("project ID is required")
	}

	if p.Name == "" {
		return errors.New("project name cannot be empty")
	}

	return nil
}
