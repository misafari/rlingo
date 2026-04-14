package postgresql

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	db "github.com/misafari/rlingo/internal/db/generated"
	projectDomain "github.com/misafari/rlingo/internal/project/domain"
)

type RepositoryImplementation struct {
	queries *db.Queries
	pool    *pgxpool.Pool
}

func (r *RepositoryImplementation) Create(ctx context.Context, project *projectDomain.Project) (*projectDomain.Project, error) {
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

	id, err := qtx.CreateProject(ctx, db.CreateProjectParams{
		ID:          project.ID,
		TenantID:    project.TenantID,
		Name:        project.Name,
		Description: project.Description,
		Status:      db.ProjectStatus(project.Status),
		CreatedBy:   project.CreatedBy,
		CreatedAt:   project.CreatedAt,
		UpdatedAt:   project.UpdatedAt,
	})

	if err != nil {
		return nil, err
	}

	if err = tx.Commit(ctx); err != nil {
		return nil, err
	}

	project.ID = id

	return project, nil
}

func (r *RepositoryImplementation) Update(ctx context.Context, project *projectDomain.Project) error {
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

	fetchedProject, err := qtx.GetProjectByIDAndTenantID(ctx, db.GetProjectByIDAndTenantIDParams{
		ID:       project.ID,
		TenantID: project.TenantID,
	})
	if err != nil {
		return err
	}

	_, err = qtx.UpdateProject(ctx, db.UpdateProjectParams{
		ID:          fetchedProject.ID,
		TenantID:    fetchedProject.TenantID,
		Name:        project.Name,
		Description: project.Description,
		Status:      db.ProjectStatus(project.Status),
		UpdatedAt:   project.UpdatedAt,
	})

	if err != nil {
		return err
	}

	return tx.Commit(ctx)
}

func (r *RepositoryImplementation) DeleteOneById(ctx context.Context, id uuid.UUID, tenantID uuid.UUID) error {
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

	err = qtx.DeleteProjectByIDAndTenantID(ctx, db.DeleteProjectByIDAndTenantIDParams{
		ID:       id,
		TenantID: tenantID,
	})

	if err != nil {
		return err
	}

	return tx.Commit(ctx)
}

func (r *RepositoryImplementation) FetchAll(ctx context.Context, tenantID uuid.UUID) ([]*projectDomain.Project, error) {
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

	fp, err := qtx.ListProjectsByTenantID(ctx, tenantID)
	if err != nil {
		return nil, err
	}

	projectList := make([]*projectDomain.Project, len(fp))

	for i, project := range fp {
		projectList[i] = &projectDomain.Project{
			ID:          project.ID,
			Name:        project.Name,
			Description: project.Description,
			Status:      projectDomain.ProjectStatus(project.Status),
			TenantID:    project.TenantID,
			CreatedBy:   project.CreatedBy,
			CreatedAt:   project.CreatedAt,
			UpdatedBy:   uuid.UUID{},
			UpdatedAt:   project.UpdatedAt,
		}
	}

	return projectList, nil
}

func (r *RepositoryImplementation) FetchOneByID(ctx context.Context, projectID, tenantID uuid.UUID) (*projectDomain.Project, error) {
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

	fp, err := qtx.GetProjectByIDAndTenantID(ctx, db.GetProjectByIDAndTenantIDParams{
		ID:       projectID,
		TenantID: tenantID,
	})

	if err != nil {
		return nil, err
	}

	entity := &projectDomain.Project{
		ID:          fp.ID,
		Name:        fp.Name,
		Description: fp.Description,
		Status:      projectDomain.ProjectStatus(fp.Status),
		TenantID:    fp.TenantID,
		CreatedBy:   fp.CreatedBy,
		CreatedAt:   fp.CreatedAt,
		UpdatedAt:   fp.UpdatedAt,
	}

	return entity, nil
}
