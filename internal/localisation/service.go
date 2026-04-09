package localisation

import (
	"context"

	"github.com/misafari/rlingo/internal/localisation/domain"
	"golang.org/x/text/language"
)

type Service struct {
	repository *Repository
}

func (s *Service) GetOrCreateLanguage(ctx context.Context, code language.Tag) (*domain.Language, error) {

	return nil, nil
}

func NewService(repository *Repository) *Service {
	return &Service{
		repository: repository,
	}
}
