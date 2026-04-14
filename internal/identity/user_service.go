package identity

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/misafari/rlingo/internal/identity/domain"
	"github.com/misafari/rlingo/internal/identity/errors"
)

type UserService interface {
	SignUp(ctx context.Context, user *domain.User) (*SignUpResult, error)
}

type userServiceImpl struct {
	repository   UserRepository
	tokenService TokenService
}

type SignUpResult struct {
	UserID      uuid.UUID
	Email       string
	FullName    string
	AccessToken string
	ExpiresAt   time.Time
}

func (u *userServiceImpl) SignUp(ctx context.Context, user *domain.User) (*SignUpResult, error) {
	user.Email = strings.ToLower(strings.TrimSpace(user.Email))
	user.FullName = strings.TrimSpace(user.FullName)

	if err := user.ValidateSignUpInput(); err != nil {
		return nil, err
	}

	if err := user.HashPassword(user.PasswordHash); err != nil {
		return nil, err
	}

	taken, err := u.repository.EmailExists(ctx, user.Email)
	if err != nil {
		return nil, fmt.Errorf("SignUp: check email: %w", err)
	}
	if taken {
		return nil, errors.ErrEmailTaken
	}

	user.Status = domain.UsersStatusACTIVE
	newUser, err := u.repository.CreateNewUser(ctx, user)
	if err != nil {
		return nil, fmt.Errorf("SignUp: create new user: %w", err)
	}

	fmt.Println(newUser.ID)

	token, ext, err := u.tokenService.IssueToken(newUser.ID)
	if err != nil {
		return nil, fmt.Errorf("SignUp: issue token: %w", err)
	}

	return &SignUpResult{
		UserID:      newUser.ID,
		Email:       newUser.Email,
		FullName:    newUser.FullName,
		AccessToken: token,
		ExpiresAt:   ext,
	}, nil
}

func NewUserService(repository UserRepository, tokenService TokenService) UserService {
	if repository == nil {
		panic("repository is nil")
	}

	if tokenService == nil {
		panic("tokenService is nil")
	}

	return &userServiceImpl{
		repository:   repository,
		tokenService: tokenService,
	}
}
