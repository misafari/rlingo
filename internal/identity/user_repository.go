package identity

import (
	"context"

	"github.com/misafari/rlingo/internal/identity/domain"
)

type UserRepository interface {
	EmailExists(ctx context.Context, email string) (bool, error)
	CreateNewUser(ctx context.Context, user *domain.User) (*domain.User, error)
}
