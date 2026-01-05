package postgres

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/misafari/rlingo/internal/domain/locale"
)

type LocaleRepository struct {
	pool *pgxpool.Pool
}

func (l *LocaleRepository) Create(ctx context.Context, locale *locale.Locale) error {
	tx, err := l.pool.Begin(ctx)
	if err != nil {
		return err
	}

	defer func() {
		if err != nil {
			_ = tx.Rollback(ctx)
		}
	}()

	query := `INSERT INTO locale (id, project_id, locale, is_default) VALUES ($1, $2, $3, $4) RETURNING id`

	err = tx.QueryRow(ctx, query, uuid.New(), locale.ProjectID, locale.Locale, locale.IsDefault).Scan(&locale.ID)
	if err != nil {
		return err
	}

	return tx.Commit(ctx)
}

func (l *LocaleRepository) FetchAll(ctx context.Context) ([]*locale.Locale, error) {
	query := `SELECT id, project_id, locale, is_default FROM locale`
	rows, err := l.pool.Query(ctx, query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var entities []*locale.Locale

	for rows.Next() {
		var entity locale.Locale
		if err = rows.Scan(&entity.ID, &entity.ProjectID, entity.Locale, entity.IsDefault); err != nil {
			return nil, err
		}

		entities = append(entities, &entity)
	}

	return entities, nil
}

func (l *LocaleRepository) DeleteOneById(ctx context.Context, id uuid.UUID) error {
	tx, err := l.pool.Begin(ctx)
	if err != nil {
		return err
	}

	defer func() {
		if err != nil {
			_ = tx.Rollback(ctx)
		}
	}()

	q := "DELETE from locale where id = $1"
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

func (l *LocaleRepository) Update(ctx context.Context, entity *locale.Locale) error {
	tx, err := l.pool.Begin(ctx)
	if err != nil {
		return err
	}

	defer func() {
		if err != nil {
			_ = tx.Rollback(ctx)
		}
	}()

	q := "UPDATE locale SET project_id = $1, locale = $2, is_default = $3 WHERE id = $4"
	tag, err := tx.Exec(ctx, q, entity.ProjectID, entity.Locale, entity.IsDefault, entity.ID)
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

func NewLocalRepository(pool *pgxpool.Pool) locale.Repository {
	return &LocaleRepository{
		pool: pool,
	}
}
