package translation

import (
	"time"

	"github.com/google/uuid"
)

type Translation struct {
	ID        uuid.UUID
	Key       string
	Locale    string
	Text      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
