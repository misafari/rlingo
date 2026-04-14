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
	Signup(ctx context.Context, user *domain.User) (*SignInResult, error)
	Signin(ctx context.Context, email, pass string) (*SignInResult, error)
}

type userServiceImpl struct {
	repository   UserRepository
	tokenService TokenService
}

type SignInResult struct {
	UserID      uuid.UUID
	Email       string
	FullName    string
	AccessToken string
	ExpiresAt   time.Time
}

func (u *userServiceImpl) Signup(ctx context.Context, user *domain.User) (*SignInResult, error) {
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

	token, ext, err := u.tokenService.IssueToken(newUser.ID)
	if err != nil {
		return nil, fmt.Errorf("SignUp: issue token: %w", err)
	}

	return &SignInResult{
		UserID:      newUser.ID,
		Email:       newUser.Email,
		FullName:    newUser.FullName,
		AccessToken: token,
		ExpiresAt:   ext,
	}, nil
}

func (u *userServiceImpl) Signin(ctx context.Context, email, pass string) (*SignInResult, error) {
	email = strings.ToLower(strings.TrimSpace(email))
	pass = strings.TrimSpace(pass)

	if email == "" || pass == "" {
		return nil, fmt.Errorf("SignIn: email or password is empty")
	}

	fetchedUser, err := u.repository.FindByEmail(ctx, email)
	if err != nil {
		return nil, fmt.Errorf("SignIn: find by email: %w", err)
	}

	if fetchedUser == nil {
		return nil, fmt.Errorf("SignIn: user not found")
	}

	err = fetchedUser.CheckPassword(pass)
	if err != nil {
		return nil, fmt.Errorf("SignIn: check password failed: %w", err)
	}

	token, ext, err := u.tokenService.IssueToken(fetchedUser.ID)
	if err != nil {
		return nil, fmt.Errorf("SignIn: issue token: %w", err)
	}

	return &SignInResult{
		UserID:      fetchedUser.ID,
		Email:       fetchedUser.Email,
		FullName:    fetchedUser.FullName,
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
