package postgres

import (
	"context"
	"fmt"
	"log"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	domain "github.com/misafari/rlingo/internal/domain/translation_key"
)

type TranslationKeyRepository struct {
	pool *pgxpool.Pool
}

func (t *TranslationKeyRepository) DeleteOneById(ctx context.Context, id uuid.UUID) error {
	tx, err := t.pool.Begin(ctx)
	if err != nil {
		return err
	}

	defer func(tx pgx.Tx, ctx context.Context) {
		if err = tx.Rollback(ctx); err != nil {
			log.Printf("error rolling back transaction: %v", err)
		}
	}(tx, ctx)

	query := `DELETE FROM translation_key WHERE id = $1`
	tag, err := tx.Exec(ctx, query, id)
	if err != nil {
		return err
	}

	rows := tag.RowsAffected()

	if rows == 0 {
		return pgx.ErrNoRows
	}

	if rows > 1 {
		return pgx.ErrTooManyRows
	}

	return tx.Commit(ctx)
}

func (t *TranslationKeyRepository) Update(ctx context.Context, entity *domain.Key) error {
	tx, err := t.pool.Begin(ctx)
	if err != nil {
		return err
	}

	defer func(tx pgx.Tx, ctx context.Context) {
		if err = tx.Rollback(ctx); err != nil {
			log.Printf("error rolling back transaction key: %v", err)
		}
	}(tx, ctx)

	query := `UPDATE translation_key SET project_id = $1, key = $2 WHERE id = $3`

	tag, err := tx.Exec(ctx, query, entity.ProjectID, entity.Key, entity.ID)

	if err != nil {
		return fmt.Errorf("error updating translation key: %w", err)
	}

	if tag.RowsAffected() == 0 {
		return pgx.ErrNoRows
	}

	if tag.RowsAffected() > 1 {
		return pgx.ErrTooManyRows
	}

	if err = tx.Commit(ctx); err != nil {
		return fmt.Errorf("error committing transaction key: %w", err)
	}

	return nil
}

func (t *TranslationKeyRepository) FetchAll(ctx context.Context) ([]*domain.Key, error) {
	query := `SELECT id, project_id, key FROM translation_key`
	rows, err := t.pool.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("error fetching translations key: %w", err)
	}
	defer rows.Close()

	var keys []*domain.Key

	for rows.Next() {
		var k domain.Key
		if err = rows.Scan(&k.ID, &k.ProjectID, &k.Key); err != nil {
			return nil, fmt.Errorf("error scanning translation key: %w", err)
		}

		keys = append(keys, &k)
	}

	return keys, nil
}

func (t *TranslationKeyRepository) Create(ctx context.Context, entity *domain.Key) error {
	tx, err := t.pool.Begin(ctx)
	if err != nil {
		return fmt.Errorf("error starting transaction key: %w", err)
	}
	defer func() {
		if err != nil {
			_ = tx.Rollback(ctx)
		}
	}()

	query := `INSERT INTO translation_key (id, project_id, key) VALUES ($1, $2, $3) RETURNING id`
	err = tx.QueryRow(ctx, query, uuid.New(), entity.ProjectID, entity.Key).
		Scan(&entity.ID)

	if err != nil {
		return fmt.Errorf("error inserting translation key: %w", err)
	}

	if err = tx.Commit(ctx); err != nil {
		return fmt.Errorf("error committing transaction key: %w", err)
	}

	return nil
}

func NewTranslationKeyRepository(pool *pgxpool.Pool) domain.Repository {
	return &TranslationKeyRepository{
		pool: pool,
	}
}
