package project

import (
	"context"
	"github.com/misafari/rlingo/internal/db/generated"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository struct {
	queries *generated.Queries
	pool *pgxpool.Pool
}

func (r *Repository) Create(ctx context.Context, project *project.Project) error {
	tx, err := r.pool.Begin(ctx)
    if err != nil {
        return err
    }

    defer func() {
        if err != nil {
            _ = tx.Rollback(ctx)
        }
    }()

	qtx := r.queries.WithTx(tx)

	_, err = qtx.CreateProject(ctx, generated.CreateProjectParams{
		ID: project.ID,
		TenantID: project.TenantID,
		Name:     project.Name,
		Description: project.Description,
		Status: generated.ProjectStatus(project.Status),
		CreatedBy:   project.CreatedBy,
		CreatedAt:   project.CreatedAt,
		UpdatedAt:   project.UpdatedAt,
	})

	if err != nil {
		return err
	}

	return tx.Commit(ctx)
}

func (r *Repository) DeleteOneById(ctx context.Context, id uuid.UUID, tenantID uuid.UUID) error {
	tx, err := r.pool.Begin(ctx)
	if err != nil {
		return err
	}

	defer func() {
		if err != nil {
			_ = tx.Rollback(ctx)
		}
	}()

	qtx := r.queries.WithTx(tx)

	_, err = qtx.DeleteProjectByIDAndTenantID(ctx, generated.DeleteProjectByIDAndTenantIDParams{
		ID: id,
		TenantID: tenantID,
	})

	if err != nil {
		return err
	}

	return tx.Commit(ctx)
}

func (r *Repository) FetchAll(ctx context.Context, tenantID uuid.UUID) ([]*project.Project, error) {
	tx, err := r.pool.Begin(ctx)
	if err != nil {
		return nil, err
	}

	defer func() {
		if err != nil {
			_ = tx.Rollback(ctx)
		}
	}()

	qtx := r.queries.WithTx(tx)

	projects, err = qtx.ListProjectsByTenantID(ctx, tenantID)
	if err != nil {
		return nil, err
	}

	return projects, nil
}