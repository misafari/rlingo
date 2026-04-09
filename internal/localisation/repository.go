package localisation

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	db "github.com/misafari/rlingo/internal/db/generated"
	"github.com/misafari/rlingo/internal/localisation/domain"
)

type Repository struct {
	queries *db.Queries
	pool    *pgxpool.Pool
}

func (r *Repository) CreateIfNotExists(ctx context.Context, language *domain.Language) (*domain.Language, error) {
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

	existsByCode, err := qtx.CheckIfLanguageExistsByCode(ctx, language.Code.String())
	if err != nil {
		return nil, err
	}

	if existsByCode {
		return language, nil
	}

	_, err = qtx.CreateLanguage(ctx, db.CreateLanguageParams{
		Code:       language.Code.String(),
		Name:       language.Name,
		NativeName: language.NativeName,
		Rtl:        language.IsRtl,
	})

	if err != nil {
		return nil, err
	}

	if err = tx.Commit(ctx); err != nil {
		return nil, err
	}

	return language, nil
}
