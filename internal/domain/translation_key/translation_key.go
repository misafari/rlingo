package translation_key

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type Key struct {
	ID        uuid.UUID
	ProjectID uuid.UUID
	Key       string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func Validate(k *Key, idCheck bool) error {
	if k == nil {
		return errors.New("translation key is nil")
	}

	if idCheck && k.ID == uuid.Nil {
		return errors.New("translation key ID is required")
	}

	if k.ProjectID == uuid.Nil {
		return errors.New("translation key project id cannot be empty")
	}

	if k.Key == "" {
		return errors.New("translation key cannot be empty")
	}

	return nil
}
