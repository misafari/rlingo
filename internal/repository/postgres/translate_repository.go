package postgres

import (
	"context"
	"fmt"
	"log"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	domain "github.com/misafari/rlingo/internal/domain/translation"
)

type TranslationRepository struct {
	pool *pgxpool.Pool
}

func (t *TranslationRepository) DeleteOneById(ctx context.Context, id uuid.UUID) error {
	tx, err := t.pool.Begin(ctx)
	if err != nil {
		return err
	}

	defer func(tx pgx.Tx, ctx context.Context) {
		if err = tx.Rollback(ctx); err != nil {
			log.Printf("error rolling back transaction: %v", err)
		}
	}(tx, ctx)

	query := `DELETE FROM translation WHERE id = $1`
	tag, err := tx.Exec(ctx, query, id.String())
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

func (t *TranslationRepository) Update(ctx context.Context, translation *domain.Translation) error {
	tx, err := t.pool.Begin(ctx)
	if err != nil {
		return err
	}

	defer func(tx pgx.Tx, ctx context.Context) {
		if err = tx.Rollback(ctx); err != nil {
			log.Printf("error rolling back transaction: %v", err)
		}
	}(tx, ctx)

	query := `UPDATE translation SET key_id = $1, locale_id = $2, text = $3 WHERE id = $4`

	tag, err := tx.Exec(ctx, query, translation.KeyID, translation.LocaleID, translation.Text, translation.ID)

	if err != nil {
		return fmt.Errorf("error updating translation: %w", err)
	}

	if tag.RowsAffected() == 0 {
		return pgx.ErrNoRows
	}

	if tag.RowsAffected() > 1 {
		return pgx.ErrTooManyRows
	}

	if err = tx.Commit(ctx); err != nil {
		return fmt.Errorf("error committing transaction: %w", err)
	}

	return nil
}

func (t *TranslationRepository) FetchAll(ctx context.Context) ([]*domain.Translation, error) {
	query := `SELECT id, key_id, locale_id, text FROM translation`

	rows, err := t.pool.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("error fetching translations: %w", err)
	}
	defer rows.Close()

	var translations []*domain.Translation

	for rows.Next() {
		var translation domain.Translation
		if err = rows.Scan(&translation.ID, &translation.KeyID, &translation.LocaleID, &translation.Text); err != nil {
			return nil, fmt.Errorf("error scanning translation: %w", err)
		}

		translations = append(translations, &translation)
	}

	return translations, nil
}

func (t *TranslationRepository) Create(ctx context.Context, translation *domain.Translation) error {
	tx, err := t.pool.Begin(ctx)
	if err != nil {
		return fmt.Errorf("error starting transaction: %w", err)
	}

	defer func() {
		if err != nil {
			_ = tx.Rollback(ctx)
		}
	}()

	query := `INSERT INTO translation (id, key_id, locale_id, text) VALUES ($1, $2, $3, $4) RETURNING id`
	err = tx.QueryRow(ctx, query, uuid.New(), translation.KeyID, translation.LocaleID, translation.Text).
		Scan(&translation.ID)

	if err != nil {
		return fmt.Errorf("error inserting translation: %w", err)
	}

	if err = tx.Commit(ctx); err != nil {
		return fmt.Errorf("error committing transaction: %w", err)
	}

	return nil
}

func NewTranslationRepository(pool *pgxpool.Pool) domain.Repository {
	return &TranslationRepository{
		pool: pool,
	}
}
