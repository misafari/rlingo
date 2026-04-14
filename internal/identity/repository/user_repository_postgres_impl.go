package repository

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
	db "github.com/misafari/rlingo/internal/db/generated"
	"github.com/misafari/rlingo/internal/identity"
	"github.com/misafari/rlingo/internal/identity/domain"
)

type userRepositoryPostgresImpl struct {
	queries *db.Queries
	pool    *pgxpool.Pool
}

func (u *userRepositoryPostgresImpl) EmailExists(ctx context.Context, email string) (bool, error) {
	tx, err := u.pool.Begin(ctx)
	if err != nil {
		return false, err
	}

	defer func() {
		if err != nil {
			_ = tx.Rollback(ctx)
		}
	}()

	qtx := u.queries.WithTx(tx)

	isExist, err := qtx.CheckIfUserExistsByEmail(ctx, email)
	if err != nil {
		return false, err
	}

	if err = tx.Commit(ctx); err != nil {
		return false, err
	}

	return isExist, nil
}

func (u *userRepositoryPostgresImpl) CreateNewUser(ctx context.Context, user *domain.User) (*domain.User, error) {
	tx, err := u.pool.Begin(ctx)
	if err != nil {
		return nil, err
	}

	defer func() {
		if err != nil {
			_ = tx.Rollback(ctx)
		}
	}()

	qtx := u.queries.WithTx(tx)

	createdUserID, err := qtx.CreateUser(ctx, db.CreateUserParams{
		ID:    uuid.New(),
		Email: user.Email,
		FullName: pgtype.Text{
			String: user.FullName,
			Valid:  true,
		},
		PasswordHash: pgtype.Text{
			String: user.PasswordHash,
			Valid:  true,
		},
		IsSso:  false,
		Status: db.UsersStatusACTIVE,
	})
	if err != nil {
		return nil, err
	}

	if err = tx.Commit(ctx); err != nil {
		return nil, err
	}

	if createdUserID == uuid.Nil {
		return nil, errors.New("id is empty")
	}
	user.ID = createdUserID

	return user, nil
}

func (u *userRepositoryPostgresImpl) FindByEmail(ctx context.Context, email string) (*domain.User, error) {
	tx, err := u.pool.Begin(ctx)
	if err != nil {
		return nil, err
	}

	defer func() {
		if err != nil {
			_ = tx.Rollback(ctx)
		}
	}()

	qtx := u.queries.WithTx(tx)

	fetchedUser, err := qtx.FindUserOneByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	if err = tx.Commit(ctx); err != nil {
		return nil, err
	}

	if fetchedUser.ID == uuid.Nil {
		return nil, errors.New("user not found")
	}

	return &domain.User{
		ID:           fetchedUser.ID,
		Email:        fetchedUser.Email,
		FullName:     fetchedUser.FullName.String,
		IsSso:        fetchedUser.IsSso,
		PasswordHash: fetchedUser.PasswordHash.String,
		Status:       domain.UsersStatus(fetchedUser.Status),
		LastLoginAt:  fetchedUser.LastLoginAt,
		CreatedAt:    fetchedUser.CreatedAt,
	}, nil
}

func NewUserRepositoryPostgresImpl(
	queries *db.Queries,
	pool *pgxpool.Pool,
) identity.UserRepository {
	if queries == nil {
		panic("queries is required")
	}

	if pool == nil {
		panic("pool is required")
	}

	return &userRepositoryPostgresImpl{
		queries: queries,
		pool:    pool,
	}
}
