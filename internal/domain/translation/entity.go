package translation

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type Status string

const (
	StatusDraft    Status = "DRAFT"
	StatusReviewed Status = "REVIEWED"
	StatusApproved Status = "APPROVED"
)

type Translation struct {
	ID        uuid.UUID `json:"id"`
	TenantID  uuid.UUID `json:"tenant_id"`
	ProjectID uuid.UUID `json:"project_id"`
	Key       string    `json:"key"`
	Locale    string    `json:"locale"`
	Text      string    `json:"text"`
	Status    Status    `json:"status"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (t *Translation) CanEdit() error {
	if t.Status == StatusApproved {
		return errors.New("cannot edit an approved translation")
	}
	return nil
}
