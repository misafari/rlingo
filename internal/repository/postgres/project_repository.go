package postgres

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/misafari/rlingo/internal/domain/project"
)

type ProjectRepository struct {
	pool *pgxpool.Pool
}

func (p *ProjectRepository) Update(ctx context.Context, entity *project.Project) error {
	tx, err := p.pool.Begin(ctx)
	if err != nil {
		return err
	}

	defer func() {
		if err != nil {
			_ = tx.Rollback(ctx)
		}
	}()

	q := "UPDATE project SET name = $1 WHERE id = $2"
	tag, err := tx.Exec(ctx, q, entity.Name, entity.ID)
	if err != nil {
		return err
	}

	if tag.RowsAffected() == 0 {
		return pgx.ErrNoRows
	}

	if tag.RowsAffected() > 1 {
		return pgx.ErrTooManyRows
	}

	return tx.Commit(ctx)
}

func (p *ProjectRepository) DeleteOneById(ctx context.Context, id uuid.UUID) error {
	tx, err := p.pool.Begin(ctx)
	if err != nil {
		return err
	}

	defer func() {
		if err != nil {
			_ = tx.Rollback(ctx)
		}
	}()

	q := "DELETE from project where id = $1"
	tag, err := tx.Exec(ctx, q, id)
	if err != nil {
		return err
	}

	if tag.RowsAffected() == 0 {
		return pgx.ErrNoRows
	}

	if tag.RowsAffected() > 1 {
		return pgx.ErrTooManyRows
	}

	return tx.Commit(ctx)
}

func (p *ProjectRepository) Create(ctx context.Context, project *project.Project) error {
	tx, err := p.pool.Begin(ctx)
	if err != nil {
		return err
	}

	defer func() {
		if err != nil {
			_ = tx.Rollback(ctx)
		}
	}()

	query := `INSERT INTO project (id, name) VALUES ($1, $2) RETURNING id`
	err = tx.QueryRow(ctx, query, uuid.New(), project.Name).Scan(&project.ID)
	if err != nil {
		return err
	}

	return tx.Commit(ctx)
}

func (p *ProjectRepository) FetchAll(ctx context.Context) ([]*project.Project, error) {
	query := `SELECT id, name FROM project`
	rows, err := p.pool.Query(ctx, query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var projects []*project.Project

	for rows.Next() {
		var p project.Project
		if err = rows.Scan(&p.ID, &p.Name); err != nil {
			return nil, err
		}

		projects = append(projects, &p)
	}

	return projects, nil
}

func NewProjectRepository(pool *pgxpool.Pool) project.Repository {
	return &ProjectRepository{
		pool: pool,
	}
}
