package domain

import (
	"time"

	"github.com/google/uuid"
	"golang.org/x/text/language"
)

type ProjectLanguage struct {
	ProjectID    uuid.UUID
	LanguageCode language.Tag
	IsBase       bool
	EnabledAt    time.Time
}
