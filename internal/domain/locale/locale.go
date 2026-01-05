package locale

import (
	"errors"

	"github.com/google/uuid"
	"golang.org/x/text/language"
)

type Locale struct {
	ID        uuid.UUID
	ProjectID uuid.UUID
	Locale    language.Tag
	IsDefault bool
}

func Validate(l *Locale, idCheck bool) error {
	if l == nil {
		return errors.New("entity is nil")
	}

	if idCheck && l.ID == uuid.Nil {
		return errors.New("entity ID is required")
	}

	if l.ProjectID == uuid.Nil {
		return errors.New("entity project id is nil")
	}

	if l.Locale == (language.Tag{}) || l.Locale == language.Und {
		return errors.New("entity locale is empty or undetermined")
	}

	if l.Locale.String() == "und" {
		return errors.New("entity locale must be a valid ISO code")
	}

	return nil
}
