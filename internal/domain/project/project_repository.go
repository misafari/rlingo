package project

import (
	"context"

	"github.com/google/uuid"
)

type Repository interface {
	Create(context.Context, *Project) error
	FetchAll(context.Context) ([]*Project, error)
	DeleteOneById(context.Context, uuid.UUID) error
	Update(context.Context, *Project) error
}
