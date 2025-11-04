package db

import (
    "context"
    "database/sql"

    "github.com/misafari/rlingo/internal/domain/user"
)

type UserRepo struct {
    db *sql.DB
}

func NewUserRepo(db *sql.DB) *UserRepo {
    return &UserRepo{db: db}
}

func (r *UserRepo) Create(ctx context.Context, u *user.User) error {
    _, err := r.db.ExecContext(ctx,
        `INSERT INTO users (id, tenant_id, email, password_hash, name, role)
         VALUES ($1, $2, $3, $4, $5, $6)`,
        u.ID, u.TenantID, u.Email, u.PasswordHash, u.Name, u.Role,
    )
    return err
}

func (r *UserRepo) FindByEmail(ctx context.Context, email string) (*user.User, error) {
    row := r.db.QueryRowContext(ctx,
        `SELECT id, tenant_id, email, password_hash, name, role, created_at, updated_at
         FROM users WHERE email=$1`, email,
    )

    var u user.User
    if err := row.Scan(&u.ID, &u.TenantID, &u.Email, &u.PasswordHash, &u.Name, &u.Role, &u.CreatedAt, &u.UpdatedAt); err != nil {
        return nil, err
    }
    return &u, nil
}
