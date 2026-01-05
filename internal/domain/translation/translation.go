package translation

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type Translation struct {
	ID        uuid.UUID
	KeyID     uuid.UUID
	LocaleID  uuid.UUID
	Text      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func Validate(t *Translation, idCheck bool) error {
	if t == nil {
		return errors.New("translation is nil")
	}

	if idCheck && t.ID == uuid.Nil {
		return errors.New("translation ID is required")
	}

	if t.KeyID == uuid.Nil {
		return errors.New("translation key id cannot be empty")
	}

	if t.LocaleID == uuid.Nil {
		return errors.New("translation locale id cannot be empty")
	}

	if t.Text == "" {
		return errors.New("translation text cannot be empty")
	}

	return nil
}
