package postgres

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	domain "github.com/misafari/rlingo/internal/domain/translation"
)

type TranslationRepository struct {
	pool *pgxpool.Pool
}

func NewTranslationRepository(pool *pgxpool.Pool) domain.TranslationRepository {
	return &TranslationRepository{
		pool: pool,
	}
}

func (t *TranslationRepository) CreateNewTranslation(ctx context.Context, translation *domain.Translation) error {
	tx, err := t.pool.Begin(ctx)
	if err != nil {
		return fmt.Errorf("error starting transaction: %w", err)
	}
	defer func() {
		if err != nil {
			_ = tx.Rollback(ctx)
		}
	}()

	query := `INSERT INTO translation (id, key, locale, text) VALUES ($1, $2, $3, $4) RETURNING id`
	err = tx.QueryRow(ctx, query, translation.ID, translation.Key, translation.Locale, translation.Text).
		Scan(&translation.ID)

	if err != nil {
		return fmt.Errorf("error inserting translation: %w", err)
	}

	if err = tx.Commit(ctx); err != nil {
		return fmt.Errorf("error committing transaction: %w", err)
	}

	return nil
}
