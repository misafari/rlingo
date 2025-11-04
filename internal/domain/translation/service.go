package translation

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type Service struct {
	Repo Repository
}

func NewService(Repo Repository) *Service {
	return &Service{Repo: Repo}
}

func (s *Service) Create(ctx context.Context, tenantID, projectID uuid.UUID, key, locale, text string) (*Translation, error) {
	t := &Translation{
		ID:        uuid.New(),
		TenantID:  tenantID,
		ProjectID: projectID,
		Key:       key,
		Locale:    locale,
		Text:      text,
		Status:    StatusDraft,
		UpdatedAt: time.Now(),
	}
	return t, s.Repo.Create(ctx, t)
}

func (s *Service) UpdateText(ctx context.Context, tenantID uuid.UUID, id string, newText string) error {
	t, err := s.Repo.GetByID(ctx, tenantID.String(), id)
	if err != nil {
		return err
	}
	if err := t.CanEdit(); err != nil {
		return err
	}
	t.Text = newText
	t.Status = StatusReviewed
	t.UpdatedAt = time.Now()
	return s.Repo.Update(ctx, t)
}

func (s *Service) Approve(ctx context.Context, tenantID uuid.UUID, id string) error {
	t, err := s.Repo.GetByID(ctx, tenantID.String(), id)
	if err != nil {
		return err
	}
	if t.Status == StatusApproved {
		return nil
	}
	t.Status = StatusApproved
	t.UpdatedAt = time.Now()
	return s.Repo.Update(ctx, t)
}
