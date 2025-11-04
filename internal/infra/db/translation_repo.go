package db

import (
	"context"
	"database/sql"
	"time"

	"github.com/misafari/rlingo/internal/domain/translation"
)

type TranslationRepo struct {
	DB *sql.DB
}

func NewTranslationRepo(db *sql.DB) *TranslationRepo {
	return &TranslationRepo{DB: db}
}

func (r *TranslationRepo) Create(ctx context.Context, t *translation.Translation) error {
	_, err := r.DB.ExecContext(ctx,
		`INSERT INTO translations (id, tenant_id, project_id, key, locale, text, status, updated_at)
		 VALUES ($1,$2,$3,$4,$5,$6,$7,$8)`,
		t.ID, t.TenantID, t.ProjectID, t.Key, t.Locale, t.Text, t.Status, t.UpdatedAt,
	)
	return err
}

func (r *TranslationRepo) GetByID(ctx context.Context, tenantID, id string) (*translation.Translation, error) {
	row := r.DB.QueryRowContext(ctx,
		`SELECT id, tenant_id, project_id, key, locale, text, status, updated_at
		 FROM translations WHERE tenant_id=$1 AND id=$2`,
		tenantID, id,
	)
	var t translation.Translation
	err := row.Scan(&t.ID, &t.TenantID, &t.ProjectID, &t.Key, &t.Locale, &t.Text, &t.Status, &t.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func (r *TranslationRepo) Update(ctx context.Context, t *translation.Translation) error {
	_, err := r.DB.ExecContext(ctx,
		`UPDATE translations SET text=$1, status=$2, updated_at=$3 WHERE id=$4 AND tenant_id=$5`,
		t.Text, t.Status, time.Now(), t.ID, t.TenantID,
	)
	return err
}

func (r *TranslationRepo) Delete(ctx context.Context, tenantID, id string) error {
	_, err := r.DB.ExecContext(ctx,
		`DELETE FROM translations WHERE tenant_id=$1 AND id=$2`, tenantID, id)
	return err
}

func (r *TranslationRepo) ListByProject(ctx context.Context, tenantID, projectID string) ([]*translation.Translation, error) {
	rows, err := r.DB.QueryContext(ctx,
		`SELECT id, tenant_id, project_id, key, locale, text, status, updated_at
		 FROM translations WHERE tenant_id=$1 AND project_id=$2`, tenantID, projectID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var out []*translation.Translation
	for rows.Next() {
		var t translation.Translation
		if err := rows.Scan(&t.ID, &t.TenantID, &t.ProjectID, &t.Key, &t.Locale, &t.Text, &t.Status, &t.UpdatedAt); err != nil {
			return nil, err
		}
		out = append(out, &t)
	}
	return out, nil
}
